// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/yguilai/pipiao-bot/app/transport/internal/conf"
	"github.com/yguilai/pipiao-bot/app/transport/internal/data"
	"github.com/yguilai/pipiao-bot/app/transport/internal/handler"
	"github.com/yguilai/pipiao-bot/app/transport/internal/server"
	"github.com/yguilai/pipiao-bot/app/transport/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	client := data.NewAsynqClient(confData)
	messageForwardService := service.NewMessageForwardService(logger, confData, client)
	eventHandler := handler.NewEventHandler(messageForwardService)
	botServer := server.NewBotServer(logger, eventHandler, confData)
	app := newApp(logger, botServer)
	return app, func() {
	}, nil
}