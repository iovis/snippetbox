package models

import (
	"database/sql"
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
	return Snippet{}, nil
}

func (self *SnippetModel) Latest() ([]Snippet, error) {
	return nil, nil
}
