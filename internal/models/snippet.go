package models

import (
	"database/sql"
	"errors"
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
	DB *sql.DB
}

func (self *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	result, err := self.DB.Exec(`
    insert into
      snippets (title, content, created, expires)
    values
      (
        ?,
        ?,
        utc_timestamp(),
        date_add(utc_timestamp(), interval ? day)
      );
    `,
		title,
		content,
		expires,
	)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (self *SnippetModel) Get(id int) (Snippet, error) {
	row := self.DB.QueryRow(`
    select
      id,
      title,
      content,
      created,
      expires
    from
      snippets
    where
      expires > utc_timestamp()
      and id = ?;
    `,
		id,
	)

	var snippet Snippet
	err := row.Scan(
		&snippet.ID,
		&snippet.Title,
		&snippet.Content,
		&snippet.Created,
		&snippet.Expires,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Snippet{}, ErrNoRecord
		} else {
			return Snippet{}, err
		}
	}

	return snippet, nil
}

func (self *SnippetModel) Latest() ([]Snippet, error) {
	return nil, nil
}
