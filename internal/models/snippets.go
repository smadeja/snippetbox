package models

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *pgxpool.Pool
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := "insert into snippets (title, content, created, expires) " +
		"values ($1, $2, current_timestamp at time zone 'utc', current_timestamp at time zone 'utc' + interval '1 day' * $3) " +
		"returning id"

	var id int
	err := m.DB.QueryRow(context.Background(), stmt, title, content, expires).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return []*Snippet{}, nil
}
