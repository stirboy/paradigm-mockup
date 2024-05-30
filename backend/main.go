package main

import (
	"backend/app"
	"backend/app/config"
	"backend/app/config/logger"
	"context"
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/signal"
)

func main() {
	c, err := config.NewConfig()
	if err != nil {
		fmt.Println("failed to load config: ", err)
		return
	}

	l, err := logger.New(c.DevMode)
	if err != nil {
		fmt.Println("failed to create logger: ", err)
		return
	}
	defer func(l *zap.Logger) {
		err := l.Sync()
		if err != nil {
			fmt.Println("failed to sync logger: ", err)
		}
	}(l)
	a, err := app.New(c)
	if err != nil {
		fmt.Println("failed to create app: ", err)
		return
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err = a.Start(ctx)
	if err != nil {
		fmt.Println("failed to start server: ", err)
		return
	}
}
