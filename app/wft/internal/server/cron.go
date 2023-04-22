package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/hibiken/asynq"
	"github.com/yguilai/pipiao-bot/app/wft/internal/conf"
	"github.com/yguilai/pipiao-bot/app/wft/internal/csts"
	"github.com/yguilai/pipiao-bot/pkg/consts"
)

type CronTaskServer struct {
	log       *log.Helper
	scheduler *asynq.Scheduler
	cc        *conf.Refresh
}

func NewCronTaskServer(lg log.Logger, scheduler *asynq.Scheduler, cc *conf.Refresh) *CronTaskServer {
	return &CronTaskServer{
		log:       log.NewHelper(log.With(lg, consts.ModuleKey, "wft/server/cron")),
		scheduler: scheduler,
		cc:        cc,
	}
}

func (s *CronTaskServer) Start(_ context.Context) error {
	wmRefreshTask := asynq.NewTask(csts.WmRefreshType, nil, asynq.TaskID(csts.WmRefreshType))
	wfRefreshTask := asynq.NewTask(csts.WfOfficalType, nil, asynq.TaskID(csts.WfOfficalType))
	_, _ = s.scheduler.Register(
		s.cc.WmCron,
		wmRefreshTask,
		asynq.Queue(csts.CronQueue),
	)
	_, _ = s.scheduler.Register(
		s.cc.WfCron,
		wfRefreshTask,
		asynq.Queue(csts.CronQueue),
	)
	return s.scheduler.Start()
}

func (s *CronTaskServer) Stop(_ context.Context) error {
	s.scheduler.Shutdown()
	return nil
}
