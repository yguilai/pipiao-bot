package handler

import (
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/event"
	"github.com/yguilai/pipiao-bot/app/transport/internal/logic"
	"github.com/yguilai/pipiao-bot/app/transport/internal/svc"
)

func newAtMessageHandler(svcCtx *svc.ServiceContext) event.ATMessageEventHandler {
	messageLogic := logic.NewMessageLogic(svcCtx)
	return func(event *dto.WSPayload, data *dto.WSATMessageData) error {
		return messageLogic.Handle(event, (*dto.Message)(data))
	}
}
