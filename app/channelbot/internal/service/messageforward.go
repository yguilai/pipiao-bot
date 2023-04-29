package service

import (
	"context"

	"github.com/bytedance/sonic"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/hibiken/asynq"
	"github.com/tencent-connect/botgo/dto"
	"github.com/yguilai/pipiao-bot/app/channelbot/internal/conf"
	"github.com/yguilai/pipiao-bot/pkg/consts"
)

type MessageForwardService struct {
	logger *log.Helper
	c      *conf.Data
	ac     *asynq.Client
}

func NewMessageForwardService(logger log.Logger, c *conf.Data, ac *asynq.Client) *MessageForwardService {
	return &MessageForwardService{
		logger: log.NewHelper(log.With(logger, consts.ModuleKey, "transport.service.message_forward")),
		c:      c,
		ac:     ac,
	}
}

func (s *MessageForwardService) Handle(_ *dto.WSPayload, data *dto.Message) error {
	msgBytes, err := sonic.Marshal(data)
	if err != nil {
		return err
	}
	task := asynq.NewTask(s.c.Asynq.Type, msgBytes)
	_, err = s.ac.EnqueueContext(context.Background(), task, asynq.Queue(s.c.Asynq.Queue))
	if err != nil {
		s.logger.Error(err)
	}
	return err
}
