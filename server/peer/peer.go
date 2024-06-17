package peer

import (
	"bytes"
	"net"
)

type Peer struct {
	Con net.Conn

	msgCh chan *Message
	delCh chan *Peer
}

func NewPeer(con net.Conn, delCh chan *Peer, msgCh chan *Message) *Peer {
	return &Peer{
		Con:   con,
		delCh: delCh,
		msgCh: msgCh,
	}
}

func (p *Peer) Send(msg []byte) (int, error) {
	return p.Con.Write(append(msg, '\n'))
}

func (p *Peer) OK() (int, error) {
	return p.Send([]byte(OK))
}

func (p *Peer) Reader() error {
	buf := make([]byte, 1024)

	for {
		n, err := p.Con.Read(buf)

		if err != nil {
			p.delCh <- p
			break
		}

		data := buf[:n]

		fields := bytes.Fields(data)

		if len(fields) == 0 {
			p.Send([]byte("ERR invalid command with 0 len"))
			continue
		}

		cmd, err := NewCommand(fields)

		if err != nil {
			p.Send([]byte(err.Error()))
			continue
		}

		p.msgCh <- &Message{
			Cmd:  cmd,
			Peer: p,
		}
	}

	return nil
}
