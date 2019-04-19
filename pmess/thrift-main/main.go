package main

import (
	"crypto/tls"
	"github.com/apache/thrift/lib/go/thrift"
	"jspring.top/pushmess/bmess"
	"jspring.top/pushmess/config"
	"jspring.top/pushmess/log"
	"jspring.top/pushmess/pmess"
	th "jspring.top/pushmess/thrift"
)

func main() {
	config.LoadConfig()
	pmess.Handle()
	AddInterruptHandler(func() {
		log.Log.Warn("Stopping Mess handle...")
		close(bmess.Quit)
		log.Log.Info("Mess handle shutdown")
	})
	protocolFactory := thrift.NewTCompactProtocolFactory()
	transportFactory := thrift.NewTBufferedTransportFactory(8192)

	go func() {
		log.Log.Infof("Experimental RPC server listening on %s",
			config.Cfg.Listen)
		err := runServer(transportFactory, protocolFactory,
			config.Cfg.Listen, false)
		log.Log.Tracef("Finished serving expimental RPC: %v",
			err)
	}()

	<-InterruptHandlersDone
	log.Log.Info("Shutdown complete")
}

func runServer(transportFactory thrift.TTransportFactory,
	protocolFactory thrift.TProtocolFactory,
	addr string, secure bool) error {
	var transport thrift.TServerTransport
	var err error
	if secure {
		cfg := new(tls.Config)
		if cert, err := tls.LoadX509KeyPair("server.crt", "server.key"); err == nil {
			cfg.Certificates = append(cfg.Certificates, cert)
		} else {
			return err
		}
		transport, err = thrift.NewTSSLServerSocket(addr, cfg)
	} else {
		transport, err = thrift.NewTServerSocket(addr)
	}
	if err != nil {
		return err
	}
	log.Log.Printf("%T\n", transport)
	handler := &pmess.PmessHandler{}
	processor := th.NewPmessServiceProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	log.Log.Println("Starting the simple server... on ", addr)
	return server.Serve()
}
