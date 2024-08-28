package handler

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtxKey          = "userId"
)

func (h *Handler) userIdentity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get(authorizationHeader)
		if header == "" {
			http.Error(w, "empty auth header", http.StatusUnauthorized)
			return
		}

		headerParts := strings.Split(header, " ")
		fmt.Println(headerParts)
		if len(headerParts) != 1 {
			http.Error(w, "invalid auth header", http.StatusUnauthorized)
			return
		}

		userId, err := h.services.Authorization.ParseToken(headerParts[0])
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), userCtxKey, userId)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
func getUserId(r *http.Request) (int, error) {
	id, ok := r.Context().Value(userCtxKey).(int)
	if !ok {
		return 0, errors.New("user id not found in context")
	}
	return id, nil
}
