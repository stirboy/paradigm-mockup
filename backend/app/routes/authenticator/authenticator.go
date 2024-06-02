package authenticator

import (
	"backend/app/database/dberrors"
	"backend/app/routes/handler"
	"backend/app/routes/routeerrors"
	"errors"
	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"go.uber.org/zap"
	"net/http"
)

func Authenticator(ja *jwtauth.JWTAuth, userHandler *handler.UserHandler) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		hfn := func(w http.ResponseWriter, r *http.Request) {
			token, claims, err := jwtauth.FromContext(r.Context())
			if err != nil {
				zap.L().Error("cant get claims from context", zap.Error(err))
				routeerrors.HandleError(w, routeerrors.Unauthorized())
				return
			}

			if token == nil || jwt.Validate(token, ja.ValidateOptions()...) != nil {
				routeerrors.HandleError(w, routeerrors.Unauthorized())
				return
			}

			userId, err := uuid.Parse(claims["user_id"].(string))
			if err != nil {
				routeerrors.HandleError(w, routeerrors.BadRequest(err.Error()))
				return
			}

			_, err = userHandler.Repo.GetUserById(r.Context(), userId)
			if err != nil {
				if errors.Is(err, dberrors.NotFoundError) {
					zap.L().Error("user not found", zap.Error(err))
					routeerrors.HandleError(w, routeerrors.Forbidden())
					return
				} else {
					zap.L().Error("cant get user by id", zap.Error(err))
					routeerrors.HandleError(w, routeerrors.BadRequest(err.Error()))
					return
				}
			}

			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(hfn)
	}
}
