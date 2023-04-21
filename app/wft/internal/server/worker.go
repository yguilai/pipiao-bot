package server

import (
	"context"
	"github.com/hibiken/asynq"
	"github.com/yguilai/pipiao-bot/app/wft/internal/csts"
	"github.com/yguilai/pipiao-bot/app/wft/internal/service"
)

type WorkerServer struct {
	server *asynq.Server
	wmr    *service.WorkerService
}

func NewWorkerServer(server *asynq.Server, wmr *service.WorkerService) *WorkerServer {
	return &WorkerServer{
		server: server,
		wmr:    wmr,
	}
}

func (s *WorkerServer) Start(_ context.Context) error {
	mux := asynq.NewServeMux()
	mux.HandleFunc(csts.WmRefreshType, s.wmr.OnWarframeMarketRefresh)
	mux.HandleFunc(csts.WfOfficalType, s.wmr.OnWarframeOfficalRefresh)
	return s.server.Start(mux)
}

func (s *WorkerServer) Stop(_ context.Context) error {
	s.server.Stop()
	return nil
}
