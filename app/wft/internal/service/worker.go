package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/hibiken/asynq"
)

type WorkerService struct {
	log *log.Helper
}

func NewWorkerService(lg log.Logger) *WorkerService {
	return &WorkerService{
		log: log.NewHelper(lg),
	}
}

func (s *WorkerService) OnWarframeMarketRefresh(_ context.Context, _ *asynq.Task) error {
	s.log.Info("OnWarframeMarketRefresh")
	return nil
}

func (s *WorkerService) OnWarframeOfficalRefresh(_ context.Context, _ *asynq.Task) error {
	s.log.Info("OnWarframeOfficalRefresh")
	return nil
}
