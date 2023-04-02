package handler

import (
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/websocket"
	"github.com/yguilai/pipiao-bot/app/channelbot/internal/handler/event"
	"github.com/yguilai/pipiao-bot/app/channelbot/internal/svc"
)

func RegisterEventHandler(svcCtx *svc.ServiceContext) dto.Intent {
	return websocket.RegisterHandlers(
		event.NewAtMessageHandler(svcCtx),
		event.NewDirectMessageHandler(svcCtx),
	)
}
