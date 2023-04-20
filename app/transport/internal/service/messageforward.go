package service

import (
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/nsqio/go-nsq"
	"github.com/tencent-connect/botgo/dto"
	"github.com/yguilai/pipiao-bot/app/transport/internal/conf"
	"github.com/yguilai/pipiao-bot/pkg/consts"
)

type MessageForwardService struct {
	logger   *log.Helper
	producer *nsq.Producer
	c        *conf.Data
}

func NewMessageForwardService(logger log.Logger, c *conf.Data) *MessageForwardService {
	producer, err := nsq.NewProducer(c.Nsq.Addr, nsq.NewConfig())
	if err != nil {
		panic(err)
	}
	return &MessageForwardService{
		logger:   log.NewHelper(log.With(logger, consts.ModuleKey, "transport.service.messageforward")),
		c:        c,
		producer: producer,
	}
}

func (s *MessageForwardService) Handle(_ *dto.WSPayload, data *dto.Message) error {
	msgBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = s.producer.Publish(s.c.Nsq.Topic, msgBytes)
	if err != nil {
		s.logger.Error(err)
	}
	return err
}
