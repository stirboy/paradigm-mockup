package authenticator

import (
	"backend/handler"
	"net/http"

	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

func Authenticator(userHandler *handler.UserHandler) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		hfn := func(w http.ResponseWriter, r *http.Request) {

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

			_, err = userHandler.Repo.GetUserById(r.Context(), userId)
			if err != nil {
				if err == pgx.ErrNoRows {
					http.Error(w, "user not found", http.StatusForbidden)
					return
				} else {
					http.Error(w, err.Error(), http.StatusBadRequest)
					return
				}
			}

			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(hfn)
	}
}
