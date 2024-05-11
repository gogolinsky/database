package database

import (
	"context"
	"database/internal/database/compute"
	"errors"
	"fmt"
	"go.uber.org/zap"
)

type computeLayer interface {
	HandleQuery(context.Context, string) (compute.Query, error)
}

type storageLayer interface {
	Get(context.Context, string) (string, bool)
	Set(context.Context, string, string) error
	Del(context.Context, string) error
}

type Database struct {
	computeLayer computeLayer
	storageLayer storageLayer
	logger       *zap.Logger
}

func NewDatabase(computeLayer computeLayer, storageLayer storageLayer, logger *zap.Logger) (*Database, error) {
	return &Database{computeLayer: computeLayer, storageLayer: storageLayer, logger: logger}, nil
}

func (database *Database) HandleQuery(ctx context.Context, rawQuery string) (string, error) {
	query, err := database.computeLayer.HandleQuery(ctx, rawQuery)
	if err != nil {
		return "", err
	}

	switch query.CommandId() {
	case compute.SetCommandId:
		return "", database.handleSetQuery(ctx, query)
	case compute.GetCommandId:
		return database.handleGetQuery(ctx, query)
	case compute.DelCommandId:
		return "", database.handleDelQuery(ctx, query)
	}

	return "", errors.New("handle query error")
}

func (database *Database) handleSetQuery(ctx context.Context, query compute.Query) error {
	arguments := query.Arguments()
	key := arguments[0]
	value := arguments[1]

	err := database.storageLayer.Set(ctx, key, value)
	if err != nil {
		return err
	}

	database.logger.Debug(fmt.Sprintf("[ok] %s %s", arguments[0], arguments[1]))

	return nil
}

func (database *Database) handleGetQuery(ctx context.Context, query compute.Query) (string, error) {
	arguments := query.Arguments()
	key := arguments[0]

	value, found := database.storageLayer.Get(ctx, key)
	if !found {
		return "", errors.New("key not found")
	}

	database.logger.Debug(fmt.Sprintf("[get ok] %s", key))

	return value, nil
}

func (database *Database) handleDelQuery(ctx context.Context, query compute.Query) error {
	arguments := query.Arguments()
	key := arguments[0]

	err := database.storageLayer.Del(ctx, key)
	if err != nil {
		return err
	}

	database.logger.Debug(fmt.Sprintf("[del ok] %s", key))

	return nil
}
