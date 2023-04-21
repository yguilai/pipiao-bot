package data

import (
	"github.com/google/wire"
	"github.com/hibiken/asynq"
	"github.com/yguilai/pipiao-bot/app/transport/internal/conf"
)

var ProviderSet = wire.NewSet(NewAsynqClient)

func NewAsynqClient(conf *conf.Data) *asynq.Client {
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr: conf.Asynq.Addr,
	})
	return client
}
