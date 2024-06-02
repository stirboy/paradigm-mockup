package handler

import (
	"backend/app/routes/authenticator/cookie"
	"backend/app/routes/routeerrors"
	"backend/app/routes/utils"
	"backend/domain/user"
	"backend/domain/user/model"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	Repo      *user.Repository
	TokenAuth *jwtauth.JWTAuth
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userId, err := utils.UserIdFromRequest(r)
	if err != nil {
		routeerrors.HandleError(w, err)
		return
	}

	u, err := h.Repo.GetUserById(r.Context(), userId)
	if err != nil {
		routeerrors.HandleError(w, routeerrors.NotFound(err.Error()))
		return
	}

	res, err := json.Marshal(u)
	if err != nil {
		routeerrors.HandleError(w, err)
		return
	}

	_, err = w.Write(res)
	if err != nil {
		routeerrors.HandleError(w, err)
		return
	}
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		routeerrors.HandleError(w, routeerrors.BadRequest(err.Error()))
		return
	}

	u, err := h.Repo.GetUserByUsername(r.Context(), body.Username)
	if err != nil {
		routeerrors.HandleError(w, routeerrors.NotFound("User not found"))
		return
	}

	if !checkPassword(u.Password, body.Password) {
		routeerrors.HandleError(w, routeerrors.BadRequest("password is incorrect"))
		return
	}

	claims := map[string]interface{}{"user_id": u.ID}
	_, tokenString, err := h.TokenAuth.Encode(claims)
	if err != nil {
		routeerrors.HandleError(w, err)
		return
	}

	cookie.SetCookie(w, tokenString)
}

func (h *UserHandler) Logout(w http.ResponseWriter, r *http.Request) {

	_, err := utils.UserIdFromRequest(r)
	if err != nil {
		routeerrors.HandleError(w, err)
		return
	}

	cookie.DeleteCookie(w)
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		routeerrors.HandleError(w, routeerrors.BadRequest(err.Error()))
		return
	}

	hashPassword, err := getHashPassword(body.Password)
	if err != nil {
		routeerrors.HandleError(w, err)
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
		routeerrors.HandleError(w, routeerrors.NotFound(err.Error()))
		return
	}

	claims := map[string]interface{}{"user_id": u.ID}
	_, tokenString, err := h.TokenAuth.Encode(claims)
	if err != nil {
		routeerrors.HandleError(w, err)
		return
	}

	cookie.SetCookie(w, tokenString)
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
