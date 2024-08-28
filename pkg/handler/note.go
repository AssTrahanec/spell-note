package handler

import (
	"SpellNote"
	"encoding/json"
	"net/http"
)

// @Summary Create note
// @Security ApiKeyAuth
// @Tags notes
// @Description create note
// @ID create-note
// @Accept  json
// @Produce  json
// @Param input body SpellNote.NoteInput true "note info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {string} string "Invalid request payload"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/note [post]
func (h *Handler) createNote(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	var input SpellNote.Note

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if input.Description == "" {
		http.Error(w, "Description is required", http.StatusBadRequest)
		return
	}
	noteID, err := h.services.Create(userId, input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"id": noteID,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

type getAllNotesResponse struct {
	Data []SpellNote.Note `json:"data"`
}

// @Summary Get User Notes
// @Security ApiKeyAuth
// @Tags notes
// @Description get all user notes
// @ID get-user-note
// @Accept  json
// @Produce  json
// @Success 200 {object} getAllNotesResponse
// @Failure 400,404 {string} string "Invalid request payload"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/note [get]
func (h *Handler) getUserNotes(w http.ResponseWriter, r *http.Request) {
	userId, err := getUserId(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	notes, err := h.services.GetNotesByUserID(userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(notes) == 0 {
		http.Error(w, "No notes found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(notes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func (h *Handler) getAllNotes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("All Notes"))
}
