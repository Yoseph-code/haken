package main

import (
	"flag"
	"log"

	"github.com/Yoseph-code/haken/internal/db"
	"github.com/Yoseph-code/haken/server"
)

func main() {
	listenAddr := flag.Uint("p", uint(server.DefaultListenAddr), "server listen address")

	flag.Parse()

	store, err := db.NewSource()

	if err != nil {
		log.Fatalf("failed to create source: %v\n", err)
	}

	s := server.New(server.Config{
		ListenAddr: uint32(*listenAddr),
		DB:         store,
	})

	if err := s.Run(); err != nil {
		log.Fatalf("failed to run server: %v\n", err)
	}
}
