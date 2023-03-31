package handler

import (
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/event"
	"github.com/yguilai/pipiao-bot/app/transport/internal/logic"
	"github.com/yguilai/pipiao-bot/app/transport/internal/svc"
)

func newDirectMessageHandler(svcCtx *svc.ServiceContext) event.DirectMessageEventHandler {
	messageLogic := logic.NewMessageLogic(svcCtx)
	return func(event *dto.WSPayload, data *dto.WSDirectMessageData) error {
		return messageLogic.Handle(event, (*dto.Message)(data))
	}
}
