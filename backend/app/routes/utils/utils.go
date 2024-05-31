package utils

import (
	"backend/app/routes/routeerrors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
	"net/http"
)

func IdFromRequest(r *http.Request) (uuid.UUID, error) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		return uuid.Nil, routeerrors.BadRequest(err.Error())
	}

	return id, nil
}

func UserIdFromRequest(r *http.Request) (uuid.UUID, error) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		return uuid.Nil, routeerrors.Unauthorized()
	}

	return ParseUUID(claims["user_id"].(string))
}

func ParseUUID(s string) (uuid.UUID, error) {
	u, err := uuid.Parse(s)
	if err != nil {
		return uuid.Nil, routeerrors.BadRequest(err.Error())
	}

	return u, nil
}
