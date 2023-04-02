package svc

import (
	"github.com/nsqio/go-nsq"
	"github.com/tencent-connect/botgo/openapi"
	"github.com/yguilai/pipiao-bot/app/channelbot/internal/commands"
	"github.com/yguilai/pipiao-bot/app/channelbot/internal/config"
	"github.com/yguilai/pipiao-bot/app/core/cmdset"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config       config.Config
	MsgProducer  *nsq.Producer
	Redis        *redis.Redis
	Api          openapi.OpenAPI
	CmdContainer *cmdset.CommandContainer
}

func NewServiceContext(c config.Config, api openapi.OpenAPI) (*ServiceContext, error) {
	nsqConfig := nsq.NewConfig()
	producer, err := nsq.NewProducer(c.Nsq.Addr, nsqConfig)
	if err != nil {
		return nil, err
	}
	return &ServiceContext{
		Config:      c,
		MsgProducer: producer,
		//Redis:        redis.MustNewRedis(c.Redis),
		Api:          api,
		CmdContainer: commands.RegisterCommands(),
	}, nil
}
