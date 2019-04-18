package bmess

import (
	"jspring.top/pushmess/thrift"
)

type PushHandle interface {
	Start()
	Trans(mt *Mt)
}

// Mt 消息结构
type Mt struct {
	Ptype  int32
	Ids    []string
	Reqstr *thrift.Tip
}

var (
	// Quit 退出信号
	Quit = make(chan struct{})
)
