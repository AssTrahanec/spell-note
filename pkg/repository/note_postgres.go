package repository

import (
	"SpellNote"
	"github.com/jmoiron/sqlx"
)

type NotePostgres struct {
	db *sqlx.DB
}

func NewNotePostgres(db *sqlx.DB) *NotePostgres {
	return &NotePostgres{db: db}
}

func (r *NotePostgres) Create(userId int, note SpellNote.Note) (int, error) {
	var id int
	query := "INSERT INTO notes (user_id ,description) values ($1, $2) returning id"
	row := r.db.QueryRow(query, userId, note.Description)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *NotePostgres) GetNotesByUserID(userId int) ([]SpellNote.Note, error) {
	var notes []SpellNote.Note
	query := "SELECT * FROM notes WHERE user_id = $1"
	err := r.db.Select(&notes, query, userId)
	if err != nil {
		return nil, err
	}
	return notes, nil
}
