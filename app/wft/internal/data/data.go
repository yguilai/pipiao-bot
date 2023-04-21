package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-redis/redis/v8"
	"github.com/yguilai/pipiao-bot/app/wft/internal/conf"
	"github.com/yguilai/pipiao-bot/pkg/consts"
	"time"

	etcdr "github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewRegistry,
	NewRedisClient,
	NewWarframeMarketEntryRepo,
)

// Data .
type Data struct {
	rds redis.Cmdable
}

// NewData .
func NewData(c *conf.Data, rds redis.Cmdable, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		rds.ShutdownSave(context.Background())
	}
	return &Data{
		rds: rds,
	}, cleanup, nil
}

func NewRedisClient(c *conf.Data, lg log.Logger) redis.Cmdable {
	logs := log.NewHelper(log.With(lg, consts.ModuleKey, "wft.data.redis"))
	client := redis.NewClient(&redis.Options{
		Addr:         c.Redis.Addr,
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
		DialTimeout:  time.Second * 2,
		PoolSize:     10,
	})
	err := client.Ping(context.Background()).Err()
	if err != nil {
		logs.Fatalf("redis connect error: %v", err)
	}
	return client
}

func NewRegistry(c *conf.Registry) registry.Registrar {
	client, err := clientv3.New(clientv3.Config{
		Endpoints: []string{c.Etcd.Addr},
	})
	if err != nil {
		panic(err)
	}
	return etcdr.New(client)
}