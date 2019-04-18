package pmess

import (
	"context"
	"github.com/apache/thrift/lib/go/thrift"
	th "jspring.top/pushmess/thrift"
	"testing"
)

func TestGetuiPush(t *testing.T) {

	transport, err := thrift.NewTSocket("localhost:9000")
	if err != nil {
		log.Fatal("Error opening socket:", err)

	}
	protocolFactory := thrift.NewTCompactProtocolFactory()
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	tspt, err := transportFactory.GetTransport(transport)
	if err != nil {
		log.Fatal(err)
	}
	defer tspt.Close()
	if err := tspt.Open(); err != nil {
		log.Fatal(err)
	}
	client := th.NewPmessServiceClientFactory(tspt, protocolFactory)
	ctx := context.Background()
	err = client.Push(ctx, []*th.Devcid{{Dev: "getui", Cid: "1e2d59000f69f24f9fdd1e4dbbd6b4e0"}},
		1,
		&th.Tip{
			Title:      "测试标题",
			Text:       "测试push",
			Appcontent: "透传内容",
		})
	// resp, err := client.Push(ctx, &PmessRequest{
	// 	Oids:   "[{\"huawei\":\"0863445031572533300003481400CN01\"}]",
	// 	Ptype:  1,
	// 	Reqstr: "{\"appcontent\":\"透传内容\",\"text\":\"测试push\",\"title\":\"测试标题\"}",
	// })
	if err != nil {
		log.Error(err)
	}

	log.Info("end:")
}

/**
curl -H "Content-Type: application/json" \
https://restapi.getui.com/v1/h9aFujUzXP9nmmtPzFXy69/auth_sign \
-XPOST -d '{ "sign":"008bcb9e33b43fc113ac174906518fc7fb98716af65da40ef42e5910fd6969bc",
"timestamp":"1553932086378",
"appkey":"9Wr21kaxjw9wyzjfGdKWb2"
}'

//bcd4ff4b038998b18f6def8009e309ff1b60fad1eb1a963f70a835bce258a3e9
//e940ec561621ac7e4f55e4bcaaf4f49d

curl -H "Content-Type: application/json" \
    -H "authtoken:bcd4ff4b038998b18f6def8009e309ff1b60fad1eb1a963f70a835bce258a3e9" \
     https://restapi.getui.com/v1/h9aFujUzXP9nmmtPzFXy69/push_single \
     -XPOST -d '{
                   "message": {
                   "appkey": "9Wr21kaxjw9wyzjfGdKWb2",
                   "is_offline": true,
                   "offline_expire_time":10000000,
                   "msgtype": "notification"
                },
                "notification": {
                    "style": {
                        "type": 0,
                        "text": "测试通知内容",
                        "title": "测试通知标题"
                    },
                    "transmission_type": true,
                    "transmission_content": "payload"
                },
                "cid": "e940ec561621ac7e4f55e4bcaaf4f49d"
                "requestid": "5cac11b21"
            }'


*/
