package config

import (
	"flag"

	"github.com/Yoseph-code/haken/constants"
)

const (
	ServerPort   string = "p"
	ServerUser   string = "u"
	ServerSecret string = "s"
	ServerDB     string = "d"
)

func DefineServerFlags() {
	flag.Uint(ServerPort, constants.DefaultServerPort, "port to listen on")
	flag.String(ServerUser, constants.DefaultServerUser, "user to access")
	flag.String(ServerSecret, constants.DefaultServerSecret, "secret")
	flag.String(ServerDB, constants.DefaultServerDB, "database")
}

func ParseFlags() {
	flag.Parse()
}
