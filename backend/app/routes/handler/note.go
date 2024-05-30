package handler

import (
	"backend/api"
	"backend/app/routes/mapper"
	"backend/domain/note"
	"backend/domain/note/model"
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
	Repo *note.Repository
}

func (h *NoteHandler) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Title string `json:"title"`
		//Content  *api.Content `json:"content"`
		ParentId string `json:"parentId"`
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

	n := model.Notes{
		ID:        noteId,
		UserID:    userId,
		Title:     body.Title,
		ParentID:  parentId,
		CreatedAt: now,
		UpdatedAt: now,
	}

	id, err := h.Repo.Insert(r.Context(), n)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	zap.L().Info("note created", zap.String("json", string(res)))

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
	archivedOnly := r.URL.Query().Get("archived")
	if archivedOnly == "true" {
		options.ArchivedOnly = true
	}
	noteModels, err := h.Repo.GetAllByUserId(r.Context(), userId, &options)
	if err != nil {
		zap.L().Error("cant get notes", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	notes := make([]api.Note, 0, len(noteModels))
	for _, v := range noteModels {
		n := mapper.MapToNoteApi(v)
		notes = append(notes, n)
	}

	res, err := json.Marshal(notes)
	if err != nil {
		zap.L().Error("cant marshal notes", zap.Error(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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

func (h *NoteHandler) ArchiveNotes(w http.ResponseWriter, r *http.Request) {
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

	err = h.Repo.ArchiveNotes(r.Context(), userId, noteId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *NoteHandler) RestoreNotes(w http.ResponseWriter, r *http.Request) {
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

	err = h.Repo.RestoreNotes(r.Context(), userId, noteId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
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

	nModel, err := h.Repo.GetByIdAndUserId(r.Context(), noteId, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	n := mapper.MapToNoteApi(nModel)

	res, err := json.Marshal(n)
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
