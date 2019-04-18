package pmess

import (
	"context"
	"jspring.top/pushmess/thrift"
)

// PmessHandler struct{} thrift接口服务
type PmessHandler struct{}

// Push thrift用户请求方法
func (p *PmessHandler) Push(ctx context.Context, oids []*thrift.Devcid, ptype int32, tip *thrift.Tip) error {
	go sendMess(oids, ptype, tip)
	return nil
}
