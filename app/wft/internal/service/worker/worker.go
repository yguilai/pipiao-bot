package worker

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/hibiken/asynq"
	"github.com/meilisearch/meilisearch-go"
)

type WorkerService struct {
	log *log.Helper
	mls *meilisearch.Client
}

func NewWorkerService(lg log.Logger, mls *meilisearch.Client) *WorkerService {
	return &WorkerService{
		log: log.NewHelper(lg),
		mls: mls,
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
