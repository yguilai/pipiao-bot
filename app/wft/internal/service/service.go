package service

import (
	"github.com/google/wire"
	"github.com/yguilai/pipiao-bot/app/wft/internal/service/worker"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewWftService,
	worker.NewWorkerService,
)
