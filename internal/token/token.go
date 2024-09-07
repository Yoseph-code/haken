package token

type TokenType uint16

const (
	Invalid TokenType = iota
	Read
	Create
	Update
	Remove
	Ping
	Help
	OK
	Exit
	Quit
	Q
	Bye
)

var Keywords = map[string]TokenType{
	"READ":   Read,
	"CREATE": Create,
	"UPDATE": Update,
	"REMOVE": Remove,
	"PING":   Ping,
	"HELP":   Help,
	"OK":     OK,
	"EXIT":   Exit,
	"QUIT":   Quit,
	"Q":      Q,
	"BYE":    Bye,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := Keywords[ident]; ok {
		return tok
	}
	return Invalid
}
