package pmess

import (
	"jspring.top/pushmess/bmess"
	"jspring.top/pushmess/getui"
	"jspring.top/pushmess/huawei"
	"jspring.top/pushmess/log"
	"jspring.top/pushmess/thrift"
)

var (
	aps = map[string]bmess.PushHandle{}
)

// Handle 启动信息处理
func Handle() {
	aps["getui"] = getui.Use()
	aps["huawei"] = huawei.Use()
	for _, v := range aps {
		if v != nil {
			go v.Start()
		}
	}
}

func sendMess(oids []*thrift.Devcid, ptype int32, reqstr *thrift.Tip) {
	ms := map[string][]string{}
	for _, v := range oids {
		if v == nil {
			continue
		}
		if ms[v.Dev] == nil {
			ms[v.Dev] = []string{v.Cid}
		} else {
			ms[v.Dev] = append(ms[v.Dev], v.Cid)
		}
	}
	for k, vs := range ms {
		if p, ok := aps[k]; ok && p != nil {
			v := &bmess.Mt{Ptype: ptype, Reqstr: reqstr, Ids: vs}
			p.Trans(v)
		} else {
			log.Log.Error("orgin:", k, " 不存在")
		}
	}

}
