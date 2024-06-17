package server

import "fmt"

const (
	DefaultListenAddr uint = 7777
)

type Config struct {
	ListenAddr uint
}

func defaultConfig() Config {
	return Config{
		ListenAddr: DefaultListenAddr,
	}
}

func (c Config) Address() string {
	return fmt.Sprintf(":%d", c.ListenAddr)
}
