package repository

import (
	"SpellNote"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user SpellNote.User) (int, error) {
	var id int
	query := "INSERT INTO users (username, password_hash) values ($1, $2) returning id"
	row := r.db.QueryRow(query, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(user SpellNote.User) (SpellNote.User, error) {
	var userResult SpellNote.User
	query := "SELECT id FROM users WHERE username = $1 and password_hash = $2"
	err := r.db.Get(&userResult, query, user.Username, user.Password)
	return userResult, err
}
