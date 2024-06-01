package server

import (
	"fmt"

	"github.com/Yoseph-code/haken/internal/db"
)

const (
	DefaultListenAddr = uint32(7001)
)

type Config struct {
	ListenAddr uint32
	db.DB
}

func (c Config) Address() string {
	return fmt.Sprintf(":%d", c.ListenAddr)
}

func defaultConfig() Config {
	return Config{
		ListenAddr: DefaultListenAddr,
	}
}
