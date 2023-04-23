package server

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/hibiken/asynq"
	"github.com/yguilai/pipiao-bot/app/channelbot/internal/conf"
	"github.com/yguilai/pipiao-bot/app/channelbot/internal/service"
)

// AsynqServer is a struct that implements the transport.Server interface
type AsynqServer struct {
	server *asynq.Server
	logger *log.Helper
	c      *conf.Data
	cbs    *service.ChannelBotService
}

// NewAsynqServer creates a new AsynqServer instance
func NewAsynqServer(server *asynq.Server, c *conf.Data, cbs *service.ChannelBotService, logger log.Logger) *AsynqServer {
	return &AsynqServer{
		server: server,
		logger: log.NewHelper(logger),
		c:      c,
		cbs:    cbs,
	}
}

// Start starts the Asynq server
func (s *AsynqServer) Start(context.Context) error {
	s.logger.Infof("Starting Asynq server...")
	mux := asynq.NewServeMux()
	mux.HandleFunc(s.c.Asynq.Type, s.cbs.OnMessage)
	return s.server.Run(mux)
}

// Stop stops the Asynq server
func (s *AsynqServer) Stop(context.Context) error {
	s.logger.Infof("Stopping Asynq server...")
	s.server.Shutdown()
	return nil
}
