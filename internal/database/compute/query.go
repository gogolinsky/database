package compute

type Query struct {
	commandId int
	arguments []string
}

func NewQuery(commandId int, arguments []string) Query {
	return Query{
		commandId: commandId,
		arguments: arguments,
	}
}

func (q *Query) CommandId() int {
	return q.commandId
}

func (q *Query) Arguments() []string {
	return q.arguments
}
