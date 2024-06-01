package authenticator

import (
	"backend/app/routes/handler"
	"errors"
	"net/http"

	"go.uber.org/zap"

	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func Authenticator(userHandler *handler.UserHandler) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		hfn := func(w http.ResponseWriter, r *http.Request) {

			_, claims, err := jwtauth.FromContext(r.Context())
			if err != nil {
				zap.L().Error("cant get claims from context", zap.Error(err))
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			userId, err := uuid.Parse(claims["user_id"].(string))
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			_, err = userHandler.Repo.GetUserById(r.Context(), userId)
			if err != nil {
				if errors.Is(err, pgx.ErrNoRows) {
					zap.L().Error("user not found", zap.Error(err))
					http.Error(w, "user not found", http.StatusForbidden)
					return
				} else {
					zap.L().Error("cant get user by id", zap.Error(err))
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
			}

			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(hfn)
	}
}
