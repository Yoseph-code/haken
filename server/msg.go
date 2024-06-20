package server

import (
	"github.com/Yoseph-code/haken/server/peer"
)

func (s *Server) handleMessage(msg *peer.Message) error {
	switch t := msg.Cmd.(type) {
	case peer.PingCommand:
		if _, err := msg.Peer.SendString(t.Val); err != nil {
			return err
		}
	case peer.CreateCommand:
		_, ok := s.db.LoadFromFile(t.Key)

		if ok {
			if _, err := msg.Peer.SendString("ERR key already exists"); err != nil {
				return err
			}
		} else {
			err := s.db.InsertToFile(t.Key, t.Val)

			if err != nil {
				if _, err := msg.Peer.SendString(err.Error()); err != nil {
					return err
				}
			}
		}
	case peer.ReadCommand:
		val, ok := s.db.LoadFromFile(t.Key)

		if !ok {
			if _, err := msg.Peer.SendString("ERR key not found"); err != nil {
				return err
			}
		} else {
			if _, err := msg.Peer.SendString(val); err != nil {
				return err
			}
		}
	case peer.UpdateCommand:
		_, ok := s.db.LoadFromFile(t.Key)

		if !ok {
			if _, err := msg.Peer.SendString("ERR key not found"); err != nil {
				return err
			}
		} else {
			err := s.db.UpdateToFile(t.Key, t.Val)

			if err != nil {
				if _, err := msg.Peer.SendString(err.Error()); err != nil {
					return err
				}
			}
		}
	case peer.RemoveCommand:
		_, ok := s.db.LoadFromFile(t.Key)

		if !ok {
			if _, err := msg.Peer.SendString("ERR key not found"); err != nil {
				return err
			}
		} else {
			err := s.db.RemoveFromFile(t.Key)

			if err != nil {
				if _, err := msg.Peer.SendString(err.Error()); err != nil {
					return err
				}
			}
		}
	}

	if _, err := msg.Peer.OK(); err != nil {
		return err
	}

	return nil
}
