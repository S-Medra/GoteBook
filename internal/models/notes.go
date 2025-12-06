package models

import (
	"database/sql"
	"time"
)

type Note struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type NoteModel struct {
	DB *sql.DB
}

func (m *NoteModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO notes(title, content, created, expires)
	Values(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *NoteModel) Get(id int) (*Note, error) {
	return nil, nil
}

func (m *NoteModel) Latest() ([]*Note, error) {
	return nil, nil
}
