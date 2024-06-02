package handler

import (
	"backend/app/routes/mapper"
	"backend/app/routes/routeerrors"
	"backend/app/routes/utils"
	"backend/domain/note"
	"backend/domain/note/model"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type NoteHandler struct {
	Repo *note.Repository
}

func (h *NoteHandler) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Title    string `json:"title"`
		ParentId string `json:"parentId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		routeerrors.HandleError(w, err)
		return
	}

	now := time.Now().UTC()
	noteId := uuid.New()

	userId, err := utils.UserIdFromRequest(r)
	if err != nil {
		routeerrors.HandleError(w, err)
		return
	}

	parentId, err := parseParentId(body.ParentId)
	if err != nil {
		routeerrors.HandleError(w, err)
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
		routeerrors.HandleError(w, err)
		return
	}

	res, err := json.Marshal(id)
	if err != nil {
		routeerrors.HandleError(w, err)
		return
	}

	_, err = w.Write(res)
	if err != nil {
		routeerrors.HandleError(w, err)
		return
	}
	w.WriteHeader(http.StatusCreated)
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
	userId, err := utils.UserIdFromRequest(r)
	if err != nil {
		routeerrors.HandleError(w, err)
		return
	}

	var options note.Options
	parentId := r.URL.Query().Get("parentId")
	if parentId != "" {
		parentId, err := utils.ParseUUID(parentId)
		if err != nil {
			zap.L().Error("cant parse parentId", zap.Error(err))
			routeerrors.HandleError(w, err)
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
		routeerrors.HandleError(w, err)
		return
	}

	res, err := json.Marshal(mapper.MapToNotesApi(noteModels))
	if err != nil {
		zap.L().Error("cant marshal notes", zap.Error(err))
		routeerrors.HandleError(w, err)
		return
	}

	_, err = w.Write(res)
	if err != nil {
		routeerrors.HandleError(w, err)
		return
	}
}

func (h *NoteHandler) UpdateById(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Title   string `json:"title"`
		Content string `json:"content"`
		Icon    string `json:"icon"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	noteId, err := utils.IdFromRequest(r)
	if err != nil {
		routeerrors.HandleError(w, err)
		return
	}

	userId, err := utils.UserIdFromRequest(r)
	if err != nil {
		routeerrors.HandleError(w, err)
		return
	}

	if err = h.Repo.UpdateByUserId(r.Context(), noteId, userId, body.Content, body.Title, body.Icon); err != nil {
		routeerrors.HandleError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *NoteHandler) ArchiveNotes(w http.ResponseWriter, r *http.Request) {
	noteId, err := utils.IdFromRequest(r)
	if err != nil {
		routeerrors.HandleError(w, err)
		return
	}

	userId, err := utils.UserIdFromRequest(r)
	if err != nil {
		routeerrors.HandleError(w, err)
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
	noteId, err := utils.IdFromRequest(r)
	if err != nil {
		routeerrors.HandleError(w, err)
		return
	}

	userId, err := utils.UserIdFromRequest(r)
	if err != nil {
		routeerrors.HandleError(w, err)
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
	noteId, err := utils.IdFromRequest(r)
	if err != nil {
		routeerrors.HandleError(w, err)
		return
	}

	userId, err := utils.UserIdFromRequest(r)
	if err != nil {
		routeerrors.HandleError(w, err)
		return
	}

	nModel, err := h.Repo.GetByIdAndUserId(r.Context(), noteId, userId)
	if err != nil {
		routeerrors.HandleError(w, err)
		return
	}

	n := mapper.MapToNoteApi(nModel)

	res, err := json.Marshal(n)
	if err != nil {
		routeerrors.HandleError(w, err)
		return
	}

	_, err = w.Write(res)
	if err != nil {
		routeerrors.HandleError(w, err)
		return
	}

	fmt.Println("get note")
}

func (h *NoteHandler) DeleteById(w http.ResponseWriter, r *http.Request) {
	noteId, err := utils.IdFromRequest(r)
	if err != nil {
		routeerrors.HandleError(w, err)
		return
	}

	userId, err := utils.UserIdFromRequest(r)
	if err != nil {
		routeerrors.HandleError(w, err)
		return
	}

	err = h.Repo.Delete(r.Context(), noteId, userId)
	if err != nil {
		zap.L().Error("cant delete notes", zap.Error(err))
		routeerrors.HandleError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
