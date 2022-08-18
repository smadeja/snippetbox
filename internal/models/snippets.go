package models

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
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
	stmt := "select * from snippets " +
		"where id = $1 and expires at time zone 'utc' > transaction_timestamp();"

	row := m.DB.QueryRow(context.Background(), stmt, id)
	s := &Snippet{}
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return []*Snippet{}, nil
}
