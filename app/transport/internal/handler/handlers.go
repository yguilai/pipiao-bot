package handler

import (
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/websocket"
	"github.com/yguilai/pipiao-bot/app/transport/internal/svc"
)

func Register(svcCtx *svc.ServiceContext) dto.Intent {
	return websocket.RegisterHandlers(
		newAtMessageHandler(svcCtx),
		newDirectMessageHandler(svcCtx),
	)
}
