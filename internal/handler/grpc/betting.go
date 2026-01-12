package grpc

import (
	"context"
	"sports-betting-helper/internal/domain"
	pb "sports-betting-helper/proto/gen"
)

type BettingHandler struct {
	pb.UnimplementedBettingServiceServer
	sportsbookService domain.SportsbookService
}

func NewBettingHandler(s domain.SportsbookService) *BettingHandler {
	return &BettingHandler{
		sportsbookService: s,
	}
}

func (h *BettingHandler) GetBookies(ctx context.Context, in *pb.GetBookmakersRequest) (*pb.GetBookmakersResponse, error) {
	filter := domain.BookmakerFilter{}
	if in.Id != nil {
		filter.ID = in.Id
	}
	if in.Enabled != nil {
		filter.Enabled = in.Enabled
	}

	bookmakers, err := h.sportsbookService.GetBookmakers(filter)
	if err != nil {
		return nil, err
	}

	var pbBookmakers []*pb.Bookmaker
	for _, b := range bookmakers {
		pbBookmakers = append(pbBookmakers, &pb.Bookmaker{
			Id:      b.ID,
			Name:    b.Name,
			Enabled: b.Enabled,
		})
	}

	return &pb.GetBookmakersResponse{
		Bookmakers: pbBookmakers,
	}, nil
}
