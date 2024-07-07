package server

import (
	"fmt"

	"github.com/Yoseph-code/haken/config"
)

type Config struct {
	ListenAddr uint
}

func defaultConfig() Config {
	return Config{
		ListenAddr: config.DefaultPort,
	}
}

func (c Config) Address() string {
	return fmt.Sprintf(":%d", c.ListenAddr)
}
