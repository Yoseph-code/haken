package server

import (
	"fmt"

	"github.com/Yoseph-code/haken/server/peer"
)

func (s *Server) handleMessage(msg *peer.Message) error {
	switch t := msg.Cmd.(type) {
	case peer.PingCommand:
		if _, err := msg.Peer.Send(t.Val); err != nil {
			return err
		}
	case peer.CreateCommand:
		fmt.Println(string(t.Val))
	}

	if _, err := msg.Peer.OK(); err != nil {
		return err
	}

	return nil
}
