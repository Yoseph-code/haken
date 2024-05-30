package main

import (
	"flag"
	"log"

	"github.com/Yoseph-code/haken/server"
)

func main() {
	listenAddr := flag.Uint("p", uint(server.DefaultListenAddr), "server listen address")

	flag.Parse()

	s := server.New(server.Config{
		ListenAddr: uint32(*listenAddr),
	})

	if err := s.Run(); err != nil {
		log.Fatalf("failed to run server: %v\n", err)
	}
}
