package compute

import (
	"context"
	"errors"
	"go.uber.org/zap"
)

const (
	setQueryArgumentsNumber = 2
	getQueryArgumentsNumber = 1
	delQueryArgumentsNumber = 1
)

var queryArgumentsNumber = map[int]int{
	SetCommandId: setQueryArgumentsNumber,
	GetCommandId: getQueryArgumentsNumber,
	DelCommandId: delQueryArgumentsNumber,
}

type Analyzer struct {
	logger *zap.Logger
}

func NewAnalyzer(logger *zap.Logger) (*Analyzer, error) {
	return &Analyzer{
		logger: logger,
	}, nil
}

func (a *Analyzer) AnalyzeQuery(ctx context.Context, tokens []string) (Query, error) {
	if len(tokens) == 0 {
		a.logger.Debug("invalid query, no tokens")
		return Query{}, errors.New("invalid query")
	}

	command := tokens[0]
	commandID := CommandNameToId(command)
	if commandID == UnknownCommandId {
		a.logger.Debug(
			"invalid command",
			zap.String("command", command),
		)
		return Query{}, errors.New("invalid command")
	}

	query := NewQuery(commandID, tokens[1:])
	argumentsNumber := queryArgumentsNumber[commandID]
	if len(query.Arguments()) != argumentsNumber {
		a.logger.Debug(
			"invalid arguments for query",
			zap.Any("args", query.Arguments()),
		)
		return Query{}, errors.New("invalid arguments")
	}

	a.logger.Debug("query analyzed")

	return query, nil
}
