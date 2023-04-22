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
	s.log.Info("start refresh warframe market entries")
	items, err := getAllWmEntry()
	if err != nil {
		s.log.Errorf("fetch wm items error %s", err.Error())
		return err
	}
	index := s.mls.Index(wmMeiliIndex)

	stats, err := index.GetStats()
	idxEmpty := false
	if err != nil {
		if e, ok := err.(*meilisearch.Error); ok && e.StatusCode == 404 {
			idxEmpty = true
		} else {
			s.log.Error(err)
			return err
		}
	}
	if idxEmpty {
		_, err := index.AddDocuments(items)
		if err != nil {
			s.log.Error(err)
			return err
		}
	} else if stats != nil && stats.NumberOfDocuments != 0 {
		s.log.Info(stats)
		_, err := index.DeleteAllDocuments()
		if err != nil {
			s.log.Error(err)
			return err
		}
		_, err = index.AddDocuments(items)
		if err != nil {
			s.log.Error(err)
			return err
		}
	}
	return nil
}

func (s *WorkerService) OnWarframeOfficalRefresh(_ context.Context, _ *asynq.Task) error {
	s.log.Info("OnWarframeOfficalRefresh")
	return nil
}
