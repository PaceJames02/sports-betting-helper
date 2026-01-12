package domain

type Bookmaker struct {
	ID      string
	Name    string
	Enabled bool
}

type SportsbookService interface {
	GetBookmakers() ([]Bookmaker, error)
}
