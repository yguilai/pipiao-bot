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
	wmi, err := s.wm.GetByName(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	var wmir pb.WmItemResp
	err = copier.Copy(&wmir, wmi)
	if err != nil {
		return nil, err
	}
	return &wmir, nil
}

func (s *WftService) GetWarframeOfficalItem(ctx context.Context, req *pb.WarframeItemReq) (*pb.WarframeItemResp, error) {
	return &pb.WarframeItemResp{}, nil
}
