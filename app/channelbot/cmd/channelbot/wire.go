//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/yguilai/pipiao-bot/app/channelbot/internal/biz"
	"github.com/yguilai/pipiao-bot/app/channelbot/internal/conf"
	"github.com/yguilai/pipiao-bot/app/channelbot/internal/data"
	"github.com/yguilai/pipiao-bot/app/channelbot/internal/server"
	"github.com/yguilai/pipiao-bot/app/channelbot/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
