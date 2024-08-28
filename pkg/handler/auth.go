package handler

import (
	"SpellNote"
	"SpellNote/pkg/service"
	"encoding/json"
	"errors"
	"net/http"
)

// @Summary Registration
// @Description registration
// @Tags auth
// @Accept  json
// @Produce  json
// @Param user body SpellNote.UserInput  true  "User info"
// @Success 200 {integer} integer 1
// @Failure 400 {string} string "Invalid request payload"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/auth/register [post]
func (h *Handler) registerUser(w http.ResponseWriter, r *http.Request) {
	var input SpellNote.User
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if input.Username == "" || input.Password == "" {
		http.Error(w, "Username and Password are required", http.StatusBadRequest)
		return
	}
	id, err := h.services.CreateUser(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"id": id,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

// @Summary Login
// @Tags auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body SpellNote.UserInput true "User info"
// @Success 200 {string} string "token"
// @Failure 400,404 {string}  string "Invalid request payload"
// @Failure 500 {string}  string "Internal Server Error"
// @Router /api/auth/login [post]
func (h *Handler) loginUser(w http.ResponseWriter, r *http.Request) {
	var input SpellNote.User
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	if input.Username == "" || input.Password == "" {
		http.Error(w, "Username and Password are required", http.StatusBadRequest)
		return
	}
	token, err := h.services.GenerateToken(input)
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"token": token,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
