package domain

type Bookmaker struct {
	ID      string
	Name    string
	Enabled bool
}

type BookmakerFilter struct {
	ID      *string
	Enabled *bool
}

type BookmakerRepository interface {
	GetBookmakers(filter BookmakerFilter) ([]Bookmaker, error)
}

type SportsbookService interface {
	GetBookmakers(filter BookmakerFilter) ([]Bookmaker, error)
}
