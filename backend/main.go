package main

import (
	"backend/app"
	"context"
	"fmt"
	"go.uber.org/zap"
	"os"
	"os/signal"
)

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		fmt.Println("failed to create logger: ", err)
		return
	}
	zap.ReplaceGlobals(logger)
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			fmt.Println("failed to sync logger: ", err)
		}
	}(logger)
	a, err := app.New()
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
