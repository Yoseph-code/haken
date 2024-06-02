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
	case CreateCommand:
		if err := msg.peer.db.Set(v.Key, v.Val); err != nil {
			return fmt.Errorf("failed to set key: %v", err)
		}

		if err := resp.
			NewWriter(msg.peer.Con).
			WriteString("OK"); err != nil {
			return err
		}
	case ReadCommand:
		val, ok := msg.peer.db.Get(v.Key)
		if !ok {
			return fmt.Errorf("key not found")
		}

		if err := resp.
			NewWriter(msg.peer.Con).
			WriteString(val); err != nil {
			return err
		}
	case PingCommand:
		if err := resp.
			NewWriter(msg.peer.Con).
			WriteString(v.Val); err != nil {
			return err
		}
	}

	return nil
}
