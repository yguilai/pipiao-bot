package service

import (
	"context"
	wftv1 "github.com/yguilai/pipiao-bot/api/wft/v1"

	pb "github.com/yguilai/pipiao-bot/api/wfs/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type WfsService struct {
	pb.UnimplementedWfsServer

	wftClient wftv1.WftClient
}

func NewWfsService(wftClient wftv1.WftClient) *WfsService {
	return &WfsService{
		wftClient: wftClient,
	}
}

func (s *WfsService) GetCycles(ctx context.Context, req *emptypb.Empty) (*pb.CyclesResp, error) {
	return &pb.CyclesResp{}, nil
}

func (s *WfsService) GetAlerts(ctx context.Context, empty *emptypb.Empty) (*pb.AlertsResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *WfsService) GetSorties(ctx context.Context, empty *emptypb.Empty) (*pb.SortiesResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *WfsService) GetVoidTrader(ctx context.Context, empty *emptypb.Empty) (*pb.VoidTraderResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *WfsService) GetWarframeMarket(ctx context.Context, req *pb.WarframeMarketReq) (*pb.WarframeMarketResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *WfsService) GetWarframeMarketRiven(ctx context.Context, req *pb.WarframeMarketRivenReq) (*pb.WarframeMarketRivenResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *WfsService) GetWarframeMarketLich(ctx context.Context, req *pb.WarframeMarketLichReq) (*pb.WarframeMarketLichResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *WfsService) GetInvasions(ctx context.Context, empty *emptypb.Empty) (*pb.InvasionsResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *WfsService) GetNightwaves(ctx context.Context, empty *emptypb.Empty) (*pb.NightwavesResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *WfsService) GetKuva(ctx context.Context, empty *emptypb.Empty) (*pb.KuvaResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *WfsService) GetVoidStorm(ctx context.Context, empty *emptypb.Empty) (*pb.VoidStormResp, error) {
	//TODO implement me
	panic("implement me")
}

func (s *WfsService) GetDailyDeal(ctx context.Context, empty *emptypb.Empty) (*pb.DailyDealsResp, error) {
	//TODO implement me
	panic("implement me")
}
