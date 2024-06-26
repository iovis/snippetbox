package models

import (
	"database/sql"
	"errors"
	"snippetbox/internal/validator"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetCreateForm struct {
	Title               string     `form:"title"`
	Content             string     `form:"content"`
	Expires             int        `form:"expires"`
	validator.Validator `form:"-"` // Ignore field
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
	rows, err := self.DB.Query(`
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
    order by
      id desc
    limit
      10
  `)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var snippets []Snippet

	for rows.Next() {
		var snippet Snippet

		err := rows.Scan(
			&snippet.ID,
			&snippet.Title,
			&snippet.Content,
			&snippet.Created,
			&snippet.Expires,
		)

		if err != nil {
			return nil, err
		}

		snippets = append(snippets, snippet)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
