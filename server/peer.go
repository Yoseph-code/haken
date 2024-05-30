package server

import (
	"net"
	"strings"

	"github.com/Yoseph-code/haken/internal/keyval"
)

type Peer struct {
	Con   net.Conn
	msgCh chan *Message
	delCh chan *Peer
	kv    *keyval.KV
}

func NewPeer(con net.Conn, delCh chan *Peer, msgCh chan *Message, kv *keyval.KV) *Peer {
	return &Peer{
		Con:   con,
		delCh: delCh,
		msgCh: msgCh,
		kv:    kv,
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
		values := strings.Split(string(data), " ")

		if len(values) == 0 {
			p.Send([]byte("ERR invalid command \n"))
		} else {
			var cmd Command

			switch values[0] {
			case GET:
				command := strings.SplitN(string(data[len(values[0])+1:]), " ", 1)

				if len(command) != 1 {
					p.Send([]byte("ERR invalid command \n"))
				} else {
					cmd = GetCommand{
						Key: []byte(strings.TrimSpace(command[0])),
					}
				}
			case SET:
				command := strings.SplitN(string(data[len(values[0])+1:]), " ", 2)

				if len(command) != 2 {
					p.Send([]byte("ERR invalid command \n"))
				} else {
					cmd = SetCommand{
						Key: []byte(command[0]),
						Val: []byte(command[1]),
					}
				}
			}

			if cmd != nil {
				p.msgCh <- &Message{
					cmd:  cmd,
					peer: p,
				}
			}
		}
		copy(buf, buf[n:])
	}

	return nil
}