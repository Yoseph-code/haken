package server

import (
	"log/slog"
	"net"

	"github.com/Yoseph-code/haken/server/peer"
)

func (s *Server) acceptPeers() error {
	for {
		con, err := s.ln.Accept()

		if err != nil {
			slog.Error("accept error", "err", err)
			continue
		}

		go s.handleConn(con)
	}
}

func (s *Server) handleConn(con net.Conn) {
	p := peer.NewPeer(con, s.delPeerCh, s.msgCh)

	s.addPeerCh <- p

	if err := p.Reader(); err != nil {
		slog.Error("peer read error", "err", err, "remoteAddr", con.RemoteAddr())
	}
}

func (s *Server) listner() {
	for {
		select {
		case p := <-s.addPeerCh:
			slog.Info("peer connected", "remoteAddr", p.Con.RemoteAddr())
			s.peers[p] = true
		case p := <-s.delPeerCh:
			slog.Info("peer disconnected", "remoteAddr", p.Con.RemoteAddr())
			delete(s.peers, p)
		case <-s.quitCh:
			return
		case msg := <-s.msgCh:
			if err := s.handleMessage(msg); err != nil {
				slog.Error("raw message error", "err", err)
			}
		}
	}
}
