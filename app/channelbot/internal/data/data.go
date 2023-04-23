package data

import (
	"github.com/hibiken/asynq"
	"github.com/yguilai/pipiao-bot/app/channelbot/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewAsynqServer,
)

// Data .
type Data struct {
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}

func NewAsynqServer(c *conf.Data) *asynq.Server {
	return asynq.NewServer(asynq.RedisClientOpt{
		Addr: c.Asynq.Addr,
	}, asynq.Config{
		Queues: map[string]int{
			c.Asynq.Queue: 1,
		},
	})
}
