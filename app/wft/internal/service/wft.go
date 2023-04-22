package service

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/yguilai/pipiao-bot/app/wft/internal/data"

	pb "github.com/yguilai/pipiao-bot/api/wft/v1"
)

type WftService struct {
	pb.UnimplementedWftServer

	wm data.IWarframeMarketEntryRepo
}

func NewWftService(wm data.IWarframeMarketEntryRepo) *WftService {
	return &WftService{
		wm: wm,
	}
}

func (s *WftService) GetWarframeMarketItem(ctx context.Context, req *pb.WmItemReq) (*pb.WmItemResp, error) {
	items, err := s.wm.GetByName(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	var res []*pb.WmItemResp_WmItem
	err = copier.Copy(res, items)
	if err != nil {
		return nil, err
	}
	return &pb.WmItemResp{Items: res}, nil
}

func (s *WftService) GetWarframeOfficalItem(ctx context.Context, req *pb.WarframeItemReq) (*pb.WarframeItemResp, error) {
	return &pb.WarframeItemResp{}, nil
}
