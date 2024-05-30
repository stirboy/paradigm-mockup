package handler

import (
	"backend/domain/user"
	"backend/domain/user/model"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	tokenName = "jwt"
)

type UserHandler struct {
	Repo      *user.Repository
	TokenAuth *jwtauth.JWTAuth
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := h.Repo.GetUserByUsername(r.Context(), body.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if !checkPassword(u.Password, body.Password) {
		http.Error(w, "password is incorrect", http.StatusBadRequest)
		return
	}

	claims := map[string]interface{}{"user_id": u.ID}
	_, tokenString, err := h.TokenAuth.Encode(claims)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 24),
		SameSite: http.SameSiteStrictMode,
		Name:     tokenName,
		Value:    tokenString,
		Path:     "/",
	})
}

func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := h.Repo.GetUserByUsername(r.Context(), body.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if !checkPassword(u.Password, body.Password) {
		http.Error(w, "password is incorrect", http.StatusBadRequest)
		return
	}

	http.SetCookie(w, &http.Cookie{
		HttpOnly: true,
		MaxAge:   -1, // delete cookie
		SameSite: http.SameSiteStrictMode,
		Name:     tokenName,
		Value:    "",
		Path:     "/",
	})
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashPassword, err := getHashPassword(body.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	now := time.Now().UTC()
	u := model.Users{
		ID:        uuid.New(),
		Username:  body.Username,
		Password:  hashPassword,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err = h.Repo.InsertUser(r.Context(), u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	claims := map[string]interface{}{"user_id": u.ID}
	_, tokenString, err := h.TokenAuth.Encode(claims)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 24),
		SameSite: http.SameSiteStrictMode,
		Name:     tokenName,
		Value:    tokenString,
		Path:     "/",
	})
}

func getHashPassword(password string) (string, error) {
	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func checkPassword(hashPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}
