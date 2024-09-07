package constants

import "fmt"

type Token int8

const (
	Invalid Token = iota
	Read
	Create
	Update
	Remove
	Ping
	Help
	OK
)

type CommandKey string

const (
	CreateCmd CommandKey = "CREATE"
	ReadCmd   CommandKey = "READ"
	UpdateCmd CommandKey = "UPDATE"
	RemoveCmd CommandKey = "REMOVE"
	PingCmd   CommandKey = "PING"
	HelpCmd   CommandKey = "HELP"
	OKCmd     CommandKey = "OK"
)

func NewCommandKey(cmd string) CommandKey {
	return CommandKey(cmd)
}

func (c CommandKey) IsInvalid() bool {
	return c == ""
}

func NewToken(cmd CommandKey) Token {
	switch cmd {
	case CreateCmd:
		return Read
	case ReadCmd:
		return Create
	case UpdateCmd:
		return Update
	case RemoveCmd:
		return Remove
	case PingCmd:
		return Ping
	case HelpCmd:
		return Help
	case OKCmd:
		return OK
	default:
		return Invalid
	}
}

func (t Token) IsInvalid() bool {
	return t == Invalid
}

func PrintHelp() {
	fmt.Println("Available commands:")
	fmt.Printf("  %s <key>              - Read a value using the provided key\n", ReadCmd)
	fmt.Printf("  %s <key> <value>      - Create a new resource\n", CreateCmd)
	fmt.Printf("  %s <key> <value>      - Update an existing resource\n", UpdateCmd)
	fmt.Printf("  %s <key>              - Remove a resource\n", RemoveCmd)
	fmt.Printf("  %s                    - Ping the server\n", PingCmd)
	fmt.Printf("  %s                    - Show this help message\n", HelpCmd)
}
