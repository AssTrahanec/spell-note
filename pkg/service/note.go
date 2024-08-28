package service

import (
	"SpellNote"
	"SpellNote/pkg/repository"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type NoteService struct {
	repo repository.Note
}

func NewNoteService(repo repository.Note) *NoteService {
	return &NoteService{repo: repo}
}

func (s *NoteService) Create(userId int, note SpellNote.Note) (int, error) {
	ok, err := s.YandexSpeller(note.Description)
	if ok {
		return s.repo.Create(userId, note)
	}
	return 0, err
}

func (s *NoteService) GetNotesByUserID(userId int) ([]SpellNote.Note, error) {
	return s.repo.GetNotesByUserID(userId)
}

type SpellingError struct {
	Code int      `json:"code"`
	Pos  int      `json:"pos"`
	Row  int      `json:"row"`
	Col  int      `json:"col"`
	Len  int      `json:"len"`
	Word string   `json:"word"`
	S    []string `json:"s"`
}

func (s *NoteService) YandexSpeller(text string) (bool, error) {

	values := url.Values{}
	values.Add("text", text)
	values.Add("lang", "Russian")
	values.Add("options", "")
	values.Add("format", "")

	// Создаем URL для запроса
	spellerUrl := "https://speller.yandex.net/services/spellservice.json/checkText?" + values.Encode()
	resp, err := http.Get(spellerUrl)
	if err != nil {
		return false, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var spellingErrors []SpellingError
	if err := json.Unmarshal(body, &spellingErrors); err != nil {
		return false, fmt.Errorf("error parsing JSON response: %w", err)
	}
	fmt.Println(string(body))
	if len(spellingErrors) == 0 {
		return true, nil
	}

	var errorMessages []string
	for _, e := range spellingErrors {
		message := fmt.Sprintf("Error in word '%s': possible corrections: %v", e.Word, e.S[0])
		errorMessages = append(errorMessages, message)
	}

	// Возвращаем все сообщения об ошибках
	return false, errors.New(fmt.Sprintf("Found spelling errors:\n%s", strings.Join(errorMessages, "\n")))
}
