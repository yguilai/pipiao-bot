//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/yguilai/pipiao-bot/app/wfs/internal/conf"
	"github.com/yguilai/pipiao-bot/app/wfs/internal/data"
	"github.com/yguilai/pipiao-bot/app/wfs/internal/server"
	"github.com/yguilai/pipiao-bot/app/wfs/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(log.Logger, *conf.Server, *conf.Data, *conf.Registry) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, service.ProviderSet, newApp))
}
