package compute

const (
	UnknownCommandId = 0
	GetCommandId     = 1
	SetCommandId     = 2
	DelCommandId     = 3
)

const (
	UnknownCommand = "UNKNOWN"
	GetCommand     = "GET"
	SetCommand     = "SET"
	DelCommand     = "DEL"
)

var commandMap = map[string]int{
	UnknownCommand: UnknownCommandId,
	GetCommand:     GetCommandId,
	SetCommand:     SetCommandId,
	DelCommand:     DelCommandId,
}

func CommandNameToId(name string) int {
	value, found := commandMap[name]
	if !found {
		return UnknownCommandId
	}

	return value
}
