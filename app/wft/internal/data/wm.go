package data

import (
	"context"
	"errors"
	"github.com/bytedance/sonic"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/meilisearch/meilisearch-go"
	"github.com/yguilai/pipiao-bot/app/wft/internal/csts/mlsc"
	"github.com/yguilai/pipiao-bot/app/wft/internal/domain"
	"github.com/yguilai/pipiao-bot/pkg/consts"
)

var ErrWmEntryNotFound = errors.New("warframe market entry not found")

type IWarframeMarketEntryRepo interface {
	GetByName(ctx context.Context, name string) ([]domain.WarframeMarketItem, error)
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

func (r *wmEntryRepo) GetByName(_ context.Context, name string) ([]domain.WarframeMarketItem, error) {
	index := r.data.mls.Index(mlsc.MlsIndexWmItem)
	searchResponse, err := index.Search(name, &meilisearch.SearchRequest{
		Limit: 5,
	})
	if err != nil {
		r.log.Error(err)
		return nil, err
	}
	if searchResponse != nil && len(searchResponse.Hits) != 0 {
		var items []domain.WarframeMarketItem

		hitsBytes, err := sonic.Marshal(searchResponse.Hits)
		if err != nil {
			return nil, err
		}
		err = sonic.Unmarshal(hitsBytes, &items)
		if err != nil {
			return nil, err
		}
		return items, nil
	}
	return nil, ErrWmEntryNotFound
}
