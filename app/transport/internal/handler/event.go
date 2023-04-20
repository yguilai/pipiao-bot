package handler

import (
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/event"
	"github.com/tencent-connect/botgo/websocket"
	"github.com/yguilai/pipiao-bot/app/transport/internal/service"
	"sync"
)

type EventHandler struct {
	mfs    *service.MessageForwardService
	once   sync.Once
	indent dto.Intent
}

func NewEventHandler(mfs *service.MessageForwardService) *EventHandler {
	return &EventHandler{
		mfs: mfs,
	}
}

func (h *EventHandler) Register() dto.Intent {
	h.once.Do(func() {
		h.indent = websocket.RegisterHandlers(
			h.newAtMessageHandler(),
			h.newDirectMessageHandler(),
		)
	})
	return h.indent
}

func (h *EventHandler) newAtMessageHandler() event.ATMessageEventHandler {
	return func(e *dto.WSPayload, data *dto.WSATMessageData) error {
		return h.mfs.Handle(e, (*dto.Message)(data))
	}
}

func (h *EventHandler) newDirectMessageHandler() event.DirectMessageEventHandler {
	return func(e *dto.WSPayload, data *dto.WSDirectMessageData) error {
		return h.mfs.Handle(e, (*dto.Message)(data))
	}
}
