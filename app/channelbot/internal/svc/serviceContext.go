package svc

import (
	"github.com/nsqio/go-nsq"
	"github.com/tencent-connect/botgo/openapi"
	"github.com/yguilai/pipiao-bot/app/channelbot/internal/config"
	"github.com/yguilai/pipiao-bot/app/channelbot/internal/warframe"
)

type ServiceContext struct {
	Config      config.Config
	MsgProducer *nsq.Producer
	Api         openapi.OpenAPI
	WfListener  *warframe.WorldStateListener
}

func NewServiceContext(c config.Config, api openapi.OpenAPI) (*ServiceContext, error) {
	nsqConfig := nsq.NewConfig()
	producer, err := nsq.NewProducer(c.Nsq.Addr, nsqConfig)
	if err != nil {
		return nil, err
	}
	listener := warframe.NewListener()
	listener.Start()
	return &ServiceContext{
		Config:      c,
		MsgProducer: producer,
		Api:         api,
		WfListener:  listener,
	}, nil
}
