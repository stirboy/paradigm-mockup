package handler

import (
	"backend/model"
	"backend/model/repository/note"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
)

type NoteHandler struct {
	Repo *note.NoteRepository
}

func (h *NoteHandler) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	now := time.Now().UTC()
	noteId := uuid.New()

	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	userId, err := uuid.Parse(claims["user_id"].(string))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	note := model.Note{
		Id:        noteId,
		UserId:    userId,
		Title:     body.Title,
		Content:   body.Content,
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	if err := h.Repo.Insert(r.Context(), note); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(res)
	w.WriteHeader(http.StatusCreated)

	fmt.Println("create note")
}

func (h *NoteHandler) List(w http.ResponseWriter, r *http.Request) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	userId, err := uuid.Parse(claims["user_id"].(string))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	notes, err := h.Repo.GetAllByUserId(r.Context(), userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(notes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(res)
	w.WriteHeader(http.StatusOK)
	fmt.Println("list notes")
}

func (h *NoteHandler) UpdateById(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	noteId, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	userId, err := uuid.Parse(claims["user_id"].(string))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = h.Repo.UpdateByUserId(r.Context(), noteId, userId, body.Content, body.Title); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

	fmt.Println("update note")
}

func (h *NoteHandler) GetById(w http.ResponseWriter, r *http.Request) {
	noteId, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	userId, err := uuid.Parse(claims["user_id"].(string))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	note, err := h.Repo.GetByIdAndUserId(r.Context(), noteId, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(note)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(res)
	w.WriteHeader(http.StatusOK)

	fmt.Println("get note")
}

func (h *NoteHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete by id")
}
