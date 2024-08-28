package repository

import (
	"SpellNote"
	"github.com/jmoiron/sqlx"
)

type Note interface {
	Create(userId int, note SpellNote.Note) (int, error)
	GetNotesByUserID(userId int) ([]SpellNote.Note, error)
}

type Authorization interface {
	CreateUser(user SpellNote.User) (int, error)
	GetUser(user SpellNote.User) (SpellNote.User, error)
}

type Repository struct {
	Note
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Note:          NewNotePostgres(db),
	}
}
