package app

import (
	"backend/app/authenticator"
	"backend/app/database"
	handler2 "backend/app/handler"
	"backend/app/logger"
	"backend/app/requestid"
	"backend/domain/note"
	"backend/domain/user"
	"net/http"

	"go.uber.org/zap"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
)

func LoadRoutes(tokenAuth *jwtauth.JWTAuth, db *database.Database) *chi.Mux {
	r := chi.NewRouter()
	r.Use(requestid.RequestID)
	r.Use(logger.Logger(zap.L()))

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	userHandler := &handler2.UserHandler{
		Repo: &user.Repository{
			Db: db,
		},
		TokenAuth: tokenAuth,
	}

	noteHandler := &handler2.NoteHandler{
		Repo: &note.Repository{
			Db: db,
		},
	}

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator(tokenAuth))
		r.Use(authenticator.Authenticator(userHandler))

		r.Get("/api", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		})
		r.Route("/api/notes", loadNotesRoutes(noteHandler))
	})

	r.Group(func(r chi.Router) {
		r.Route("/api/auth", loadLoginRoutes(userHandler))
	})

	return r
}

func loadLoginRoutes(userHandler *handler2.UserHandler) func(r chi.Router) {
	return func(r chi.Router) {
		r.Post("/login", userHandler.Login)
		r.Post("/register", userHandler.Register)
		r.Post("/logout", userHandler.Logout)
	}
}

func loadNotesRoutes(noteHandler *handler2.NoteHandler) func(r chi.Router) {
	return func(r chi.Router) {
		r.Post("/", noteHandler.Create)
		r.Get("/", noteHandler.List)
		r.Get("/{id}", noteHandler.GetById)
		r.Put("/{id}", noteHandler.UpdateById)
		r.Delete("/{id}", noteHandler.DeleteById)
		r.Put("/{id}/archive", noteHandler.ArchiveNotes)
		r.Put("/{id}/restore", noteHandler.RestoreNotes)
	}

}
