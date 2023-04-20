//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/yguilai/pipiao-bot/app/wft/internal/biz"
	"github.com/yguilai/pipiao-bot/app/wft/internal/conf"
	"github.com/yguilai/pipiao-bot/app/wft/internal/server"
	"github.com/yguilai/pipiao-bot/app/wft/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
