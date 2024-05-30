package server

import "fmt"

const (
	DefaultListenAddr = uint32(7001)
)

type Config struct {
	ListenAddr uint32
}

func (c Config) Address() string {
	return fmt.Sprintf(":%d", c.ListenAddr)
}

func defaultConfig() Config {
	return Config{
		ListenAddr: DefaultListenAddr,
	}
}
