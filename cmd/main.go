package main

import (
	"bufio"
	"context"
	"database/internal/database"
	"fmt"
	"go.uber.org/zap"
	"os"
)

func main() {
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	logger.Debug("Init service")

	reader := bufio.NewReader(os.Stdin)

	db, err := database.InitDatabase(logger)
	if err != nil {
		logger.Error("failed to init database", zap.Error(err))
	}

	for {
		fmt.Print("[query] > ")

		request, err := reader.ReadString('\n')
		if err != nil {
			logger.Error("failed to read user query", zap.Error(err))
		}

		result, err := db.HandleQuery(context.Background(), request)
		if err != nil {
			logger.Error("handle query error", zap.Error(err))
		}

		fmt.Println(result)
	}
}
