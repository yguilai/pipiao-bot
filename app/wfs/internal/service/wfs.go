package service

import (
	"context"

	pb "github.com/yguilai/pipiao-bot/api/wfs/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type WfsService struct {
	pb.UnimplementedWfsServer
}

func NewWfsService() *WfsService {
	return &WfsService{}
}

func (s *WfsService) GetCycles(ctx context.Context, req *emptypb.Empty) (*pb.CyclesResp, error) {
	return &pb.CyclesResp{}, nil
}

func (s *WfsService) GetAlerts(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *WfsService) GetSorties(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *WfsService) GetVoidTrader(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *WfsService) GetWarframeMarket(ctx context.Context, req *pb.WarframeMarketReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *WfsService) GetWarframeMarketRiven(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *WfsService) GetInvasions(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *WfsService) GetNightwaves(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *WfsService) GetKuva(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *WfsService) GetVoidStorm(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *WfsService) GetDailyDeal(ctx context.Context, req *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
