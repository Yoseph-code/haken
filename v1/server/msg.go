package server

type Message struct {
	cmd  Command
	peer *Peer
}

func (s *Server) handleMessage(msg *Message) error {
	switch v := msg.cmd.(type) {
	case CreateCommand:
		err := msg.peer.db.Create(v.Key, v.Val)

		if err != nil {
			if _, err := msg.peer.Send(err.Error()); err != nil {
				return err
			}

			return err
		}

		if _, err := msg.peer.Send(OK); err != nil {
			return err
		}
	case ReadCommand:
		data, err := msg.peer.db.Read(v.Key)

		if err != nil {
			if _, err := msg.peer.Send(err.Error()); err != nil {
				return err
			}

			return err
		}

		if _, err := msg.peer.Send(data); err != nil {
			return err
		}
	case PingCommand:
		if _, err := msg.peer.Send("PONG"); err != nil {
			return err
		}
	}

	return nil
}
