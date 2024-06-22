package cli

import "fmt"

type Token int

const (
	Invalid Token = iota
	Read
	Create
	Update
	Remove
	Ping
	Help
)

func NewToken(cmd string) Token {
	switch cmd {
	case "READ":
		return Read
	case "CREATE":
		return Create
	case "UPDATE":
		return Update
	case "REMOVE":
		return Remove
	case "PING":
		return Ping
	case "HELP":
		return Help
	default:
		return Invalid
	}
}

func (t Token) IsInvalid() bool {
	return t == Invalid
}

func printHelp() {
	fmt.Println("Available commands:")
	fmt.Println("  READ <key>    - Read a value using the provided key")
	fmt.Println("  CREATE        - Create a new resource")
	fmt.Println("  UPDATE        - Update an existing resource")
	fmt.Println("  REMOVE        - Remove a resource")
	fmt.Println("  HELP          - Show this help message")
	fmt.Println("  PING          - Ping the server")
	fmt.Println("  EXIT          - Exit the application")
}
