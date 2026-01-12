package service

import (
	"sports-betting-helper/internal/domain"
)

type sportsbookService struct {
	repo domain.BookmakerRepository
}

func NewSportsbookService(repo domain.BookmakerRepository) domain.SportsbookService {
	return &sportsbookService{repo: repo}
}

func (s *sportsbookService) GetBookmakers(filter domain.BookmakerFilter) ([]domain.Bookmaker, error) {
	return s.repo.GetBookmakers(filter)
}
