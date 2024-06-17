package server

import (
	"log/slog"
	"net"

	"github.com/Yoseph-code/haken/server/peer"
)

type Server struct {
	Config

	ln net.Listener

	peers     map[*peer.Peer]bool
	addPeerCh chan *peer.Peer
	delPeerCh chan *peer.Peer
	quitCh    chan struct{}
	msgCh     chan *peer.Message
}

func New(cfg ...Config) *Server {
	c := defaultConfig()

	if len(cfg) > 0 {
		c = cfg[0]
	}

	return &Server{
		Config:    c,
		peers:     make(map[*peer.Peer]bool),
		addPeerCh: make(chan *peer.Peer),
		delPeerCh: make(chan *peer.Peer),
		quitCh:    make(chan struct{}),
		msgCh:     make(chan *peer.Message),
	}
}

func (s *Server) Run() error {
	ln, err := net.Listen("tcp", s.Config.Address())

	slog.Info("haken server is running", "listenAddr", s.ListenAddr)

	if err != nil {
		return err
	}

	s.ln = ln

	go s.listner()

	return s.acceptPeers()
}
