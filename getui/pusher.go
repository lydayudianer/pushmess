package getui

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"jspring.top/pushmess/bmess"
	"jspring.top/pushmess/config"
	"jspring.top/pushmess/thrift"
	"net/http"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

// Pusher getui推送器
type Pusher struct {
	Orgin        string
	reqkey       int64
	reqid        int64
	msch         chan *bmess.Mt
	pushurl      string
	appid        string
	appkey       string
	mastersecret string
	authtoken    string
}

var (
	auch chan struct{}
)

// Use 个推初始化
func Use() bmess.PushHandle {
	if config.Cfg.GetuiAppid == "" ||
		config.Cfg.GetuiAppkey == "" ||
		config.Cfg.GetuiAppms == "" {
		return nil
	}
	pusher := &Pusher{
		Orgin:        "getui",
		reqkey:       time.Now().Unix(),
		pushurl:      "https://restapi.getui.com/",
		appid:        config.Cfg.GetuiAppid,
		appkey:       config.Cfg.GetuiAppkey,
		mastersecret: config.Cfg.GetuiAppms,
		msch:         make(chan *bmess.Mt, 10),
	}
	log.Info("getui is started.")
	return pusher
}

func (p *Pusher) pushauth() {
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	h := sha256.New()
	h.Write([]byte(p.appkey + timestamp + p.mastersecret))
	sign := hex.EncodeToString(h.Sum(nil))
	ctx := make(map[string]string)
	ctx["sign"] = sign
	ctx["timestamp"] = timestamp
	ctx["appkey"] = p.appkey
	ctxstr, err := json.Marshal(ctx)
	if err != nil {
		log.Error(err)
		return
	}
	req, err := http.NewRequest("POST", p.pushurl+"v1/"+p.appid+"/auth_sign", strings.NewReader(string(ctxstr)))
	if err != nil {
		log.Error(err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
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
				log.Error(err)
			}
			break
		}
	}
	ret := make(map[string]string)
	err = json.Unmarshal(resb, &ret)
	if err != nil {
		log.Error(err)
		return
	}
	if ret["result"] == "ok" {
		p.authtoken = ret["auth_token"]
	} else if ret["result"] == "" {
		return
	} else {
		return
	}
	return
}

func (p *Pusher) aus() {
	timer := time.NewTicker(20 * time.Hour)
	defer timer.Stop()
	for {
		select {
		case <-bmess.Quit:
			return
		case <-auch:
			p.pushauth()
		case <-timer.C:
			p.pushauth()
		}
	}
}

func (p *Pusher) httpsend(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("authtoken", p.authtoken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
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
				log.Error(err)
			}
			break
		}
	}
	ret := make(map[string]string)
	err = json.Unmarshal(resb, &ret)
	if err != nil {
		log.Error(err)
		return
	}
	if ret["result"] == "" {
		log.Error("无返回")
	} else {
		log.Info("result:", ret["result"], " ;taskid:", ret["taskid"])
		log.Info("status:", ret["status"], " ;desc:", ret["desc"])
	}
}

func (p *Pusher) single(id string, reqstr *thrift.Tip) {
	ctx := make(map[string]interface{})
	ctx["cid"] = id
	message := make(map[string]interface{})
	message["appkey"] = p.appkey
	message["is_offline"] = true
	message["offline_expire_time"] = 172800000
	message["msgtype"] = "notification"
	ctx["message"] = message
	notification := make(map[string]interface{})
	notification["transmission_type"] = true
	style := make(map[string]interface{})
	style["type"] = 0
	notification["transmission_content"] = reqstr.Appcontent
	style["text"] = reqstr.Text
	style["title"] = reqstr.Title
	notification["style"] = style
	ctx["notification"] = notification
	ctx["requestid"] = strconv.FormatInt(p.reqkey, 16) + "gg" +
		strconv.FormatInt(atomic.AddInt64(&p.reqid, 1), 16)
	ctxstr, err := json.Marshal(ctx)
	if err != nil {
		log.Error(err)
		return
	}
	req, err := http.NewRequest("POST", p.pushurl+"v1/"+p.appid+"/push_single",
		strings.NewReader(string(ctxstr)))
	if err != nil {
		log.Error(err)
		return
	}
	p.httpsend(req)
}

func (p *Pusher) list(ids []string, reqstr string) {
	message := make(map[string]interface{})
	message["appkey"] = p.appkey
	message["is_offline"] = true
	message["offline_expire_time"] = 172800000
	message["msgtype"] = "notification"
	ctx := map[string]interface{}{
		"message": message,
	}
	notification := make(map[string]interface{})
	notification["transmission_type"] = true
	style := map[string]interface{}{
		"type": 0,
	}
	reqjson := make(map[string]string)
	err := json.Unmarshal([]byte(reqstr), &reqjson)
	if err != nil {
		log.Error("参数格式错误")
		return
	}
	notification["transmission_content"] = reqjson["appcontent"]
	style["text"] = reqjson["text"]
	style["title"] = reqjson["title"]
	notification["style"] = style
	ctx["notification"] = notification
	ctxstr, err := json.Marshal(ctx)
	if err != nil {
		log.Error(err)
		return
	}
	req, err := http.NewRequest("POST", p.pushurl+"v1/"+p.appid+"/save_list_body",
		strings.NewReader(string(ctxstr)))
	if err != nil {
		log.Error(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("authtoken", p.authtoken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
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
				log.Error(err)
			}
			break
		}
	}
	ret := make(map[string]string)
	err = json.Unmarshal(resb, &ret)
	if err != nil {
		log.Error(err)
		return
	}
	if ret["result"] != "ok" {
		log.Error("list error:", ret["result"])
		return
	}
	if ret["taskid"] == "" {
		log.Error("taskid error:", ret["desc"])
		return
	}
	ctx = map[string]interface{}{
		"taskid":      ret["taskid"],
		"cid":         ids,
		"need_detail": true,
	}
	ctxstr, err = json.Marshal(ctx)
	if err != nil {
		log.Error(err)
		return
	}
	req, err = http.NewRequest("POST", p.pushurl+"v1/"+p.appid+"/push_list",
		strings.NewReader(string(ctxstr)))
	if err != nil {
		log.Error(err)
		return
	}
	p.httpsend(req)
}

func (p *Pusher) notifyPush(mt *bmess.Mt) {
	if mt.Ids == nil || len(mt.Ids) < 1 || len(mt.Ids[0]) < 1 {
		log.Error("ids 为空")
		return
	}
	for _, v := range mt.Ids {
		p.single(v, mt.Reqstr)
	}
}

// Start 开启
func (p *Pusher) Start() {
	p.pushauth()
	go p.aus()
	for {
		select {
		case <-bmess.Quit:
			return
		case mess := <-p.msch:
			log.Info("pmess-appid:", p.appid)
			// log.Info("mess-m:", mess.m)
			if len(p.authtoken) < 1 {
				auch <- struct{}{}
				log.Error("push fail:authtoken is null")
			}
			if mess.Ptype == 1 {
				go p.notifyPush(mess)
			}

		}

	}
}

// Trans 接收代处理到消息，发送到chan
func (p *Pusher) Trans(mt *bmess.Mt) {
	p.msch <- mt
}
