package server

import (
	"net"
	"strings"

	"github.com/Yoseph-code/haken/internal/db"
)

type Peer struct {
	Con   net.Conn
	msgCh chan *Message
	delCh chan *Peer
	db    *db.DB
}

func NewPeer(con net.Conn, delCh chan *Peer, msgCh chan *Message, d *db.DB) *Peer {
	return &Peer{
		Con:   con,
		delCh: delCh,
		msgCh: msgCh,
		db:    d,
	}
}

func (p *Peer) Send(msg []byte) (int, error) {
	return p.Con.Write(msg)
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

		fields := strings.Fields(string(data))

		if len(fields) == 0 {
			p.Send([]byte("ERR invalid command \n"))
			continue
		}

		comand := fields[0]

		var cmd Command

		switch comand {
		case CREATE:
			if len(fields[1:]) < 2 {
				p.Send([]byte("ERR invalid command \n"))
				continue
			}

			key := fields[1]

			value := strings.Join(fields[2:], " ")

			cmd = CreateCommand{
				Key: key,
				Val: value,
			}
		case READ:
			if len(fields[1:]) < 1 {
				p.Send([]byte("ERR invalid command \n"))
				continue
			}

			key := fields[1]

			cmd = ReadCommand{
				Key: key,
			}
		case PING:
			cmd = PingCommand{
				Val: "PONG",
			}
		default:
			p.Send([]byte("ERR invalid command \n"))
		}

		if cmd != nil {
			p.msgCh <- &Message{
				cmd:  cmd,
				peer: p,
			}
		}

		copy(buf, buf[n:])
	}

	return nil
}
