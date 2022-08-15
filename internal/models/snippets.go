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

func (m *SnippetModel) Insert(title string, content string, expires int) (string, error) {
	stmt := "insert into snippets (title, content, created, expires) " +
		"values ($1, $2, current_timestamp at time zone 'utc', current_timestamp at time zone 'utc' + interval '$3 days')"

	result, err := m.DB.Exec(context.Background(), stmt, title, content, expires)

	if err != nil {
		return "", err
	}

	return string(result), nil
}

func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return []*Snippet{}, nil
}
