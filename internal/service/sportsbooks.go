package service

import (
	"sports-betting-helper/internal/domain"
)

type sportsbookService struct {
}

func NewSportsbookService() domain.SportsbookService {
	return &sportsbookService{}
}

func (s *sportsbookService) GetBookmakers() ([]domain.Bookmaker, error) {
	// For now, returning the same hardcoded data but from the service layer
	return []domain.Bookmaker{
		{ID: "1", Name: "Bet365", Enabled: false},
		{ID: "2", Name: "Betfair", Enabled: true},
	}, nil
}
