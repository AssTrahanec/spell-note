package service

import (
	"SpellNote"
	"SpellNote/pkg/repository"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	signingKey = "saggasgui21hui4jr12k#!"
	salt       = "sgjaglalj21oi4k"
)

var ErrUserNotFound = errors.New("user not found")

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}
func (s *AuthService) CreateUser(user SpellNote.User) (int, error) {
	user.Password = s.GeneratePasswordHash(user.Password)
	_, err := s.repo.GetUser(user)
	if err != nil {
		return s.repo.CreateUser(user)
	}
	return 0, errors.New("user already exists")
}
func (s *AuthService) GenerateToken(user SpellNote.User) (string, error) {
	user.Password = s.GeneratePasswordHash(user.Password)
	user, err := s.repo.GetUser(user)
	if err != nil {
		return "", ErrUserNotFound
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(signingKey))
}
func (s *AuthService) GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok || !token.Valid {
		return 0, errors.New("invalid token")
	}

	return claims.UserId, nil
}
