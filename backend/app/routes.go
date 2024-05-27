package app

import (
	"backend/app/authenticator"
	"backend/app/database"
	"backend/handler"
	"backend/model/note"
	"backend/model/repository/user"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
)

func LoadRoutes(tokenAuth *jwtauth.JWTAuth, db *database.Database) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
	}))

	userHandler := &handler.UserHandler{
		Repo: &user.UserRepository{
			Db: db,
		},
		TokenAuth: tokenAuth,
	}

	noteHandler := &handler.NoteHandler{
		Repo: &note.NoteRepository{
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

func loadLoginRoutes(userHandler *handler.UserHandler) func(r chi.Router) {
	return func(r chi.Router) {
		r.Post("/login", userHandler.Login)
		r.Post("/register", userHandler.Register)
		r.Post("/logout", userHandler.Logout)
	}
}

func loadNotesRoutes(noteHandler *handler.NoteHandler) func(r chi.Router) {
	return func(r chi.Router) {
		r.Post("/", noteHandler.Create)
		r.Get("/", noteHandler.List)
		r.Get("/{id}", noteHandler.GetById)
		r.Put("/{id}", noteHandler.UpdateById)
		r.Delete("/{id}", noteHandler.DeleteById)
	}

}
