package constants

type CMD string

const (
	EXIT   CMD = "exit"
	PROMPT CMD = "haken> "
	QUIT   CMD = "quit"
	Q      CMD = "q"
)

func New(cmd string) CMD {
	return CMD(cmd)
}

func (c CMD) Is(cmd string) bool {
	return c == CMD(cmd)
}

func (c CMD) IsBye() bool {
	return c.IsEmpty() || c.IsExit() || c.IsQuit()
}

func (c CMD) String() string {
	return string(c)
}

func (c CMD) IsExit() bool {
	return c == EXIT
}

func (c CMD) IsQuit() bool {
	return c == QUIT || c == Q
}

func (c CMD) IsPrompt() bool {
	return c == PROMPT
}

func (c CMD) IsEmpty() bool {
	return c == ""
}

func (c CMD) IsUnknown() bool {
	return !c.IsExit() && !c.IsQuit() && !c.IsPrompt() && !c.IsEmpty()
}

func (c CMD) IsKnown() bool {
	return !c.IsUnknown()
}
