//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/yguilai/pipiao-bot/app/transport/internal/conf"
	"github.com/yguilai/pipiao-bot/app/transport/internal/data"
	"github.com/yguilai/pipiao-bot/app/transport/internal/handler"
	"github.com/yguilai/pipiao-bot/app/transport/internal/server"
	"github.com/yguilai/pipiao-bot/app/transport/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		service.ProviderSet,
		handler.ProviderSet,
		data.ProviderSet,
		newApp,
	))
}
