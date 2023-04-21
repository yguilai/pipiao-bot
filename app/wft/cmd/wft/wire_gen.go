// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/yguilai/pipiao-bot/app/wft/internal/conf"
	"github.com/yguilai/pipiao-bot/app/wft/internal/data"
	"github.com/yguilai/pipiao-bot/app/wft/internal/server"
	"github.com/yguilai/pipiao-bot/app/wft/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, registry *conf.Registry, logger log.Logger) (*kratos.App, func(), error) {
	cmdable := data.NewRedisClient(confData, logger)
	dataData, cleanup, err := data.NewData(confData, cmdable, logger)
	if err != nil {
		return nil, nil, err
	}
	iWarframeMarketEntryRepo := data.NewWarframeMarketEntryRepo(dataData, logger)
	wftService := service.NewWftService(iWarframeMarketEntryRepo)
	grpcServer := server.NewGRPCServer(confServer, wftService, logger)
	registrar := data.NewRegistry(registry)
	app := newApp(logger, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}