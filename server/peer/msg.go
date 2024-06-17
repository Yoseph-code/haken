package peer

type Message struct {
	Cmd  Command
	Peer *Peer
}
