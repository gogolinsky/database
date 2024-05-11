package engine

import (
	"context"
	"go.uber.org/zap"
)

type Engine struct {
	memory map[string]string
	logger *zap.Logger
}

func NewEngine(
	logger *zap.Logger,
) (*Engine, error) {
	eng := new(Engine)
	eng.memory = make(map[string]string)
	eng.logger = logger

	return eng, nil
}

func (e *Engine) Set(ctx context.Context, key, value string) {
	e.memory[key] = value
}

func (e *Engine) Get(ctx context.Context, key string) (string, bool) {
	value, found := e.memory[key]
	return value, found
}

func (e *Engine) Del(ctx context.Context, key string) {
	delete(e.memory, key)
}
