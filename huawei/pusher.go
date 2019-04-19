package huawei

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"jspring.top/pushmess/bmess"
	"jspring.top/pushmess/config"
	"jspring.top/pushmess/log"
	"jspring.top/pushmess/thrift"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Pusher huawei pusher
type Pusher struct {
	Orgin            string
	msch             chan *bmess.Mt
	tokenurl         string
	pushurl          string
	appid            string
	appsec           string
	authtoken        string
	tokenExpiredTime time.Time
	authlock         sync.Mutex
	pkgName          string
}

var (
	expireDuration, _ = time.ParseDuration("48h")
)

// Use huawei 初始化
func Use() bmess.PushHandle {
	if config.Cfg.HuaweiAppid == "" ||
		config.Cfg.HuaweiAppsec == "" ||
		config.Cfg.PkgName == "" {
		return nil
	}
	v := url.Values{}
	v.Add("nsp_ctx", "{\"ver\":\"1\", \"appId\":\""+config.Cfg.HuaweiAppid+"\"}")
	pusher := &Pusher{
		Orgin:    "huawei",
		tokenurl: "https://login.cloud.huawei.com/oauth2/v2/token",
		pushurl:  "https://api.push.hicloud.com/pushsend.do?" + v.Encode(),
		appid:    config.Cfg.HuaweiAppid,
		appsec:   config.Cfg.HuaweiAppsec,
		msch:     make(chan *bmess.Mt, 10),
		pkgName:  config.Cfg.PkgName,
	}
	log.Log.Info("huawei is started")
	return pusher
}

func (p *Pusher) pushauth() {
	p.authlock.Lock()
	defer p.authlock.Unlock()
	if !p.tokenExpiredTime.IsZero() &&
		time.Now().Before(p.tokenExpiredTime) {
		return
	}
	ctx := "grant_type=client_credentials&client_secret=" + p.appsec +
		"&client_id=" + p.appid
	req, err := http.NewRequest("POST", p.tokenurl, strings.NewReader(ctx))
	if err != nil {
		log.Log.Error(err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		log.Log.Error(err)
		return
	}
	defer resp.Body.Close()
	resb := []byte{}
	for {
		buff := make([]byte, 256)
		n, err := resp.Body.Read(buff)
		if n > 0 {
			resb = append(resb, buff[:n]...)
		}
		if err != nil {
			if err.Error() != "EOF" {
				log.Log.Error(err)
			}
			break
		}
	}
	ret := make(map[string]interface{})
	err = json.Unmarshal(resb, &ret)
	if err != nil {
		log.Log.Error(err)
		return
	}
	if ret["access_token"] != "" {
		p.authtoken = fmt.Sprintf("%v", ret["access_token"])
		log.Log.Info("huawei-token:", p.authtoken)
		expiresIn := fmt.Sprintf("%v", ret["expires_in"])
		log.Log.Info("expires_in:", expiresIn)
		s, _ := time.ParseDuration(expiresIn + "s")
		p.tokenExpiredTime = time.Now().Add(s)
	}
	log.Log.Info("error:", ret["error"])
	log.Log.Info("error_description:", ret["error_description"])
}

func (p *Pusher) httpsend(req *http.Request) {
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		log.Log.Error(err)
		return
	}
	defer resp.Body.Close()
	log.Log.Info("HttpStatus:", resp.StatusCode)
	log.Log.Info("HttpHeader:", resp.Header["NSP_STATUS"])
	resb := []byte{}
	for {
		buff := make([]byte, 256)
		n, err := resp.Body.Read(buff)
		if n > 0 {
			resb = append(resb, buff[:n]...)
		}
		if err != nil {
			if err.Error() != "EOF" {
				log.Log.Error(err)
			}
			break
		}
	}
	if len(resb) < 1 {
		log.Log.Info("无返回内容")
		return
	}
	ret := make(map[string]string)
	err = json.Unmarshal(resb, &ret)
	if err != nil {
		log.Log.Error(err)
		return
	}
	log.Log.Info("code:", ret["code"], " ;msg:", ret["msg"])
	log.Log.Info("requestId:", ret["requestId"])
}

func (p *Pusher) normal(ids []string, reqstr *thrift.Tip) {
	v := url.Values{}
	v.Add("access_token", p.authtoken)
	v.Add("nsp_svc", "openpush.message.api.send")
	v.Add("nsp_ts", strconv.FormatInt(time.Now().Unix(), 10))
	idtys, err := json.Marshal(&ids)
	if err != nil {
		log.Log.Error(err)
		return
	}
	v.Add("device_token_list", string(idtys))
	v.Add("expire_time", time.Now().Add(expireDuration).Format("2006-01-02T15:04"))
	payload := map[string]interface{}{}
	hps := map[string]interface{}{}
	action := map[string]interface{}{}
	body := map[string]interface{}{}
	param := map[string]interface{}{}
	msg := map[string]interface{}{}

	body["content"] = reqstr.Text
	body["title"] = reqstr.Title

	msg["type"] = 3
	msg["body"] = body
	msg["action"] = action

	param["appPkgName"] = p.pkgName

	action["param"] = param
	action["type"] = 3

	hps["msg"] = msg

	payload["hps"] = hps

	ploadtys, err := json.Marshal(&payload)
	if err != nil {
		log.Log.Error(err)
		return
	}
	v.Add("payload", string(ploadtys))
	req, err := http.NewRequest("POST", p.pushurl,
		strings.NewReader(v.Encode()))
	if err != nil {
		log.Log.Error(err)
		return
	}
	p.httpsend(req)
}

func (p *Pusher) notifyPush(mt *bmess.Mt) {
	if mt.Ids == nil || len(mt.Ids) < 1 || len(mt.Ids[0]) < 1 {
		log.Log.Error("ids 为空")
		return
	}
	if p.tokenExpiredTime.IsZero() ||
		time.Now().After(p.tokenExpiredTime) {
		p.pushauth()
	}
	if len(mt.Ids) > 100 {
		for i, len := 0, len(mt.Ids); i < len; i += 100 {
			end := i + 100
			if end > len {
				end = len
			}
			p.normal(mt.Ids[i:end], mt.Reqstr)
		}
	} else {
		p.normal(mt.Ids, mt.Reqstr)
	}

}

// Start 启动
func (p *Pusher) Start() {
	p.pushauth()
	for {
		select {
		case <-bmess.Quit:
			return
		case mess := <-p.msch:
			log.Log.Info("huawei-appid:", p.appid)
			if mess.Ptype == 1 {
				go p.notifyPush(mess)
			}

		}

	}
}

// Trans 接收待处理的消息，发送到chan
func (p *Pusher) Trans(mt *bmess.Mt) {
	p.msch <- mt
}
