package service

import (
	"SpellNote"
	"SpellNote/pkg/repository"
)

type Authorization interface {
	CreateUser(user SpellNote.User) (int, error)
	GenerateToken(user SpellNote.User) (string, error)
	ParseToken(token string) (int, error)
}
type Note interface {
	Create(userId int, note SpellNote.Note) (int, error)
	GetNotesByUserID(userId int) ([]SpellNote.Note, error)
}
type Service struct {
	Note
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Note:          NewNoteService(repos.Note),
	}
}
