package handler

import (
	"backend/api"
	"backend/model/note"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type NoteHandler struct {
	Repo *note.NoteRepository
}

func (h *NoteHandler) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Title    string       `json:"title"`
		Content  *api.Content `json:"content"`
		ParentId string       `json:"parentId"`
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

	parentId, err := parseParentId(body.ParentId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	note := note.Note{
		Id:        &noteId,
		UserId:    &userId,
		Title:     body.Title,
		Content:   body.Content,
		ParentId:  parentId,
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

	_, err = w.Write(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)

	fmt.Println("create note")
}

func parseParentId(parentId string) (*uuid.UUID, error) {
	if parentId == "" {
		return nil, nil
	}

	id, err := uuid.Parse(parentId)
	if err != nil {
		return nil, err
	}

	return &id, nil
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

	var options note.Options
	parentId := r.URL.Query().Get("parentId")
	if parentId != "" {
		parentId, err := uuid.Parse(parentId)
		if err != nil {
			zap.L().Error("cant parse parentId", zap.Error(err))
			http.Error(w, "cant parse parentId", http.StatusBadRequest)
			return
		}
		options.ParentId = &parentId
	}
	notes, err := h.Repo.GetAllByUserId(r.Context(), userId, &options)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(notes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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

	_, err = w.Write(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("get note")
}

func (h *NoteHandler) GetByQuery(w http.ResponseWriter, r *http.Request) {
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

	_, err = w.Write(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("get note")
}

func (h *NoteHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete by id")
}
