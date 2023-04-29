package biz

import (
	"github.com/google/wire"
	"github.com/yguilai/pipiao-bot/app/channelbot/internal/biz/handler"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	handler.NewEventHandler,
)
