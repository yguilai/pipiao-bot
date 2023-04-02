package event

import (
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/event"
	"github.com/yguilai/pipiao-bot/app/channelbot/internal/logic"
	"github.com/yguilai/pipiao-bot/app/channelbot/internal/svc"
)

func NewDirectMessageHandler(svcCtx *svc.ServiceContext) event.DirectMessageEventHandler {
	messageLogic := logic.NewMessageLogic(svcCtx)
	return func(event *dto.WSPayload, data *dto.WSDirectMessageData) error {
		return messageLogic.Handle(event, (*dto.Message)(data))
	}
}
