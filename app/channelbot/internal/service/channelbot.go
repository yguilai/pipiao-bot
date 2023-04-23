package service

import (
	"context"

	"github.com/hibiken/asynq"
)

type ChannelBotService struct {
}

func NewChannelBotService() *ChannelBotService {
	return &ChannelBotService{}
}

func (s *ChannelBotService) OnMessage(_ context.Context, task *asynq.Task) error {
	return nil
}
