package compute

import (
	"context"
	"go.uber.org/zap"
)

type parser interface {
	ParseQuery(context.Context, string) ([]string, error)
}

type analyzer interface {
	AnalyzeQuery(context.Context, []string) (Query, error)
}

type Compute struct {
	parser   parser
	analyzer analyzer
	logger   *zap.Logger
}

func NewCompute(parser parser, analyzer analyzer, logger *zap.Logger) (*Compute, error) {
	return &Compute{
		parser:   parser,
		analyzer: analyzer,
		logger:   logger,
	}, nil
}

func (d *Compute) HandleQuery(ctx context.Context, rawQuery string) (Query, error) {
	tokens, err := d.parser.ParseQuery(ctx, rawQuery)
	if err != nil {
		return Query{}, err
	}

	query, err := d.analyzer.AnalyzeQuery(ctx, tokens)
	if err != nil {
		return Query{}, err
	}

	return query, nil
}
