package app

import (
	"backend/app/database"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth/v5"
)

type App struct {
	c         *Config
	r         http.Handler
	db        *database.Database
	tokenAuth *jwtauth.JWTAuth
}

func New() (*App, error) {
	c, err := NewConfig()
	if err != nil {
		fmt.Println("failed to load config: ", err)
		return nil, err
	}

	db, err := database.New(c.DbUrl)
	if err != nil {
		fmt.Println("failed to connect to database: ", err)
		return nil, err
	}

	tokenAuth := jwtauth.New("HS256", []byte(c.JwtSecret), nil)
	routes := LoadRoutes(tokenAuth, db)

	app := &App{
		c:         c,
		db:        db,
		r:         routes,
		tokenAuth: tokenAuth,
	}

	return app, nil
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":8080",
		Handler: a.r,
	}

	defer func() {
		a.db.DBPool.Close()
	}()

	ch := make(chan error, 1)

	fmt.Println("Server is running on port 8080")

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			ch <- fmt.Errorf("failed to start to server: %w", err)
		}
		close(ch)
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		return server.Shutdown(timeout)
	}
}
