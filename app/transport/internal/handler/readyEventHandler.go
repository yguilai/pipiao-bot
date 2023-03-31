package handler

import (
	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/event"
)

func newReadyEventHandler() event.ReadyHandler {
	return func(event *dto.WSPayload, data *dto.WSReadyData) {

	}
}
