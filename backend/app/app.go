package app

import (
	"backend/app/config"
	"backend/app/database"
	routes2 "backend/app/routes"
	"context"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth/v5"
)

type App struct {
	c         *config.Config
	r         http.Handler
	db        *database.Database
	tokenAuth *jwtauth.JWTAuth
}

func New(c *config.Config) (*App, error) {
	db, err := database.New(c.DbUrl)
	if err != nil {
		zap.L().Error("failed to connect to database", zap.Error(err))
		return nil, err
	}

	tokenAuth := jwtauth.New("HS256", []byte(c.JwtSecret), nil)
	routes := routes2.LoadRoutes(tokenAuth, db)

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
