package data

import (
	"context"
	etcdr "github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/middleware/circuitbreaker"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/yguilai/pipiao-bot/app/wfs/internal/conf"
	"github.com/yguilai/pipiao-bot/pkg/etcd"
	"go.opentelemetry.io/otel/propagation"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	wftv1 "github.com/yguilai/pipiao-bot/api/wft/v1"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewRegistry,
	NewWftClient,
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

func NewRegistry(c *conf.Registry) *etcdr.Registry {
	return etcd.NewEtcdRegistry([]string{c.Etcd.Addr})
}

func NewWftClient(r *etcdr.Registry) wftv1.WftClient {
	// 先不搞连接池吧
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///pipiao.wft.grpc"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
			circuitbreaker.Client(),
			tracing.Client(tracing.WithPropagator(propagation.NewCompositeTextMapPropagator(propagation.Baggage{}, propagation.TraceContext{}))),
		),
	)
	if err != nil {
		panic(err)
	}
	return wftv1.NewWftClient(conn)
}
