package server

import (
	"fmt"

	"github.com/tidwall/resp"
)

type Message struct {
	cmd  Command
	peer *Peer
}

func (s *Server) handleMessage(msg *Message) error {
	switch v := msg.cmd.(type) {
	case SetCommand:
		if err := msg.peer.db.Set(string(v.Key), v.Val); err != nil {
			return fmt.Errorf("failed to set key: %v", err)
		}

		if err := resp.
			NewWriter(msg.peer.Con).
			WriteString("OK"); err != nil {
			return err
		}
	case GetCommand:
		val, ok := msg.peer.db.Get(string(v.Key))
		if !ok {
			return fmt.Errorf("key not found")
		}
		if err := resp.
			NewWriter(msg.peer.Con).
			WriteString(fmt.Sprintf("%v", val)); err != nil {
			return err
		}
	case PingCommand:
		if err := resp.
			NewWriter(msg.peer.Con).
			WriteString(v.Value); err != nil {
			return err
		}
	}

	return nil
}
