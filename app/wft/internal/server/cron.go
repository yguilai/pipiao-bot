package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/hibiken/asynq"
	"github.com/yguilai/pipiao-bot/app/wft/internal/csts"
	"github.com/yguilai/pipiao-bot/pkg/consts"
)

type CronTaskServer struct {
	log       *log.Helper
	scheduler *asynq.Scheduler
}

func NewCronTaskServer(lg log.Logger, scheduler *asynq.Scheduler) *CronTaskServer {
	return &CronTaskServer{
		log:       log.NewHelper(log.With(lg, consts.ModuleKey, "wft/server/cron")),
		scheduler: scheduler,
	}
}

func (s *CronTaskServer) Start(_ context.Context) error {
	refreshTask := asynq.NewTask(csts.WmRefreshType, nil, asynq.TaskID(csts.WmRefreshType))
	_, _ = s.scheduler.Register(
		"@every 5s",
		refreshTask,
		asynq.Queue(csts.CronQueue),
	)
	return s.scheduler.Start()
}

func (s *CronTaskServer) Stop(_ context.Context) error {
	s.scheduler.Shutdown()
	return nil
}
