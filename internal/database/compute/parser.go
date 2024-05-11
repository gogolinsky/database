package compute

import (
	"context"
	"go.uber.org/zap"
	"strings"
)

type Parser struct {
	logger *zap.Logger
}

func NewParser(logger *zap.Logger) (*Parser, error) {
	return &Parser{
		logger: logger,
	}, nil
}

func (p *Parser) ParseQuery(ctx context.Context, query string) ([]string, error) {
	words := strings.Fields(query)

	return words, nil
}
