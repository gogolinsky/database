package database

import (
	"database/internal/database/compute"
	"database/internal/database/storage"
	"database/internal/database/storage/engine"
	"go.uber.org/zap"
)

func InitDatabase(logger *zap.Logger) (*Database, error) {
	parser, err := compute.NewParser(logger)
	if err != nil {
		return nil, err
	}

	analyzer, err := compute.NewAnalyzer(logger)
	if err != nil {
		return nil, err
	}

	computeLayer, err := compute.NewCompute(parser, analyzer, logger)
	if err != nil {
		return nil, err
	}

	eng, err := engine.NewEngine(logger)
	if err != nil {
		return nil, err
	}

	store, err := storage.NewStorage(eng, logger)
	if err != nil {
		return nil, err
	}

	db, err := NewDatabase(computeLayer, store, logger)
	if err != nil {
		return nil, err
	}

	return db, nil
}
