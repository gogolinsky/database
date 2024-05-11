package storage

import (
	"context"
	"go.uber.org/zap"
)

type Engine interface {
	Set(context.Context, string, string)
	Get(context.Context, string) (string, bool)
	Del(context.Context, string)
}

type Storage struct {
	engine Engine
	logger *zap.Logger
}

func NewStorage(
	engine Engine,
	logger *zap.Logger,
) (*Storage, error) {
	return &Storage{
		engine: engine,
		logger: logger,
	}, nil
}

func (s *Storage) Set(ctx context.Context, key, value string) error {
	s.engine.Set(ctx, key, value)
	return nil
}

func (s *Storage) Del(ctx context.Context, key string) error {
	s.engine.Del(ctx, key)
	return nil
}

func (s *Storage) Get(ctx context.Context, key string) (string, bool) {
	value, found := s.engine.Get(ctx, key)
	return value, found
}
