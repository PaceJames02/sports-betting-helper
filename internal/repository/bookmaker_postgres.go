package repository

import (
	"database/sql"
	"fmt"
	"sports-betting-helper/internal/domain"
)

type bookmakerRepository struct {
	db *sql.DB
}

func NewBookmakerRepository(db *sql.DB) domain.BookmakerRepository {
	return &bookmakerRepository{db: db}
}

func (r *bookmakerRepository) GetBookmakers(filter domain.BookmakerFilter) ([]domain.Bookmaker, error) {
	query := "SELECT id, name, enabled FROM bookmakers WHERE 1=1"
	var args []interface{}
	argCount := 1

	if filter.ID != nil {
		query += fmt.Sprintf(" AND id = $%d", argCount)
		args = append(args, *filter.ID)
		argCount++
	}

	if filter.Enabled != nil {
		query += fmt.Sprintf(" AND enabled = $%d", argCount)
		args = append(args, *filter.Enabled)
		argCount++
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookmakers []domain.Bookmaker
	for rows.Next() {
		var b domain.Bookmaker
		if err := rows.Scan(&b.ID, &b.Name, &b.Enabled); err != nil {
			return nil, err
		}
		bookmakers = append(bookmakers, b)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return bookmakers, nil
}
