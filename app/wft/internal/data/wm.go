package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/yguilai/pipiao-bot/app/wft/internal/domain"
	"github.com/yguilai/pipiao-bot/pkg/consts"
)

type IWarframeMarketEntryRepo interface {
	GetByName(ctx context.Context, name string) (*domain.WarframeMarketItem, error)
}

type wmEntryRepo struct {
	data *Data
	log  *log.Helper
}

func NewWarframeMarketEntryRepo(data *Data, logger log.Logger) IWarframeMarketEntryRepo {
	return &wmEntryRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, consts.ModuleKey, "wft/data/wm_entry")),
	}
}

func (w *wmEntryRepo) GetByName(ctx context.Context, name string) (*domain.WarframeMarketItem, error) {
	//TODO implement me
	panic("implement me")
}
