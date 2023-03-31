package svc

import (
	"github.com/nsqio/go-nsq"
	"github.com/yguilai/pipiao-bot/app/transport/internal/config"
)

type ServiceContext struct {
	Config      config.Config
	MsgProducer *nsq.Producer
}

func NewServiceContext(c config.Config) (*ServiceContext, error) {
	nsqConfig := nsq.NewConfig()
	producer, err := nsq.NewProducer(c.Nsq.Addr, nsqConfig)
	if err != nil {
		return nil, err
	}
	return &ServiceContext{
		Config:      c,
		MsgProducer: producer,
	}, nil
}
