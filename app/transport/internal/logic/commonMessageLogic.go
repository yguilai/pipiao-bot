package logic

import (
	"context"
	"encoding/json"
	"github.com/tencent-connect/botgo/dto"
	"github.com/yguilai/pipiao-bot/app/transport/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type MessageLogic struct {
	logx.Logger
	svcCtx *svc.ServiceContext
}

func NewMessageLogic(svcCtx *svc.ServiceContext) *MessageLogic {
	return &MessageLogic{
		Logger: logx.WithContext(context.Background()),
		svcCtx: svcCtx,
	}
}

func (l *MessageLogic) Handle(_ *dto.WSPayload, data *dto.Message) error {
	producer := l.svcCtx.MsgProducer
	msgBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return producer.Publish(l.svcCtx.Config.Nsq.Topic, msgBytes)
}
