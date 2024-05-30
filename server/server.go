package server

import (
	"log/slog"
	"net"

	"github.com/Yoseph-code/haken/internal/keyval"
)

type Server struct {
	Config

	ln net.Listener

	peers     map[*Peer]bool
	addPeerCh chan *Peer
	delPeerCh chan *Peer
	quitCh    chan struct{}
	msgCh     chan *Message

	kv *keyval.KV
}

func New(cfg ...Config) *Server {
	c := defaultConfig()

	if len(cfg) > 0 {
		c = cfg[0]
	}

	return &Server{
		Config:    c,
		peers:     make(map[*Peer]bool),
		addPeerCh: make(chan *Peer),
		delPeerCh: make(chan *Peer),
		quitCh:    make(chan struct{}),
		msgCh:     make(chan *Message),
		kv:        keyval.New(),
	}
}

func (s *Server) Run() error {
	ln, err := net.Listen("tcp", s.Config.Address())

	if err != nil {
		return err
	}

	s.ln = ln

	slog.Info("haken server is running", "listenAddr", s.ListenAddr)

	go s.listner()

	return s.acceptLoop()
}
