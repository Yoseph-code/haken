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
		}

		if cmd != nil {
			p.msgCh <- &Message{
				cmd:  cmd,
				peer: p,
			}
		}

		copy(buf, buf[n:])

		// if len(values) > 0 {
		// 	firstArg := values[0]
		// 	// Use the firstArg variable as needed
		// }

		// values := strings.Split(string(data), " ")
		// fmt.Println("values: ", values)
		// if len(values) == 0 {
		// 	p.Send([]byte("ERR invalid command \n"))
		// } else {
		// 	var cmd Command

		// 	switch values[0] {
		// 	case GET:
		// 		command := strings.SplitN(string(data[len(values[0])+1:]), " ", 1)

		// 		if len(command) != 1 {
		// 			p.Send([]byte("ERR invalid command \n"))
		// 		} else {
		// 			cmd = GetCommand{
		// 				Key: []byte(strings.TrimSpace(command[0])),
		// 			}
		// 		}
		// 	case SET:
		// 		command := strings.SplitN(string(data[len(values[0])+1:]), " ", 2)

		// 		if len(command) != 2 {
		// 			p.Send([]byte("ERR invalid command \n"))
		// 		} else {
		// 			cmd = SetCommand{
		// 				Key: []byte(command[0]),
		// 				Val: []byte(command[1]),
		// 			}
		// 		}
		// 	case PING:
		// 		cmd = PingCommand{
		// 			Value: "PONG",
		// 		}
		// 	}

		// 	if cmd != nil {
		// 		p.msgCh <- &Message{
		// 			cmd:  cmd,
		// 			peer: p,
		// 		}
		// 	}
		// }
	}

	return nil
}
