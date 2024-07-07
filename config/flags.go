package config

import "flag"

const (
	Server string = "s"
	Port   string = "p"
	User   string = "u"
	Secret string = ""
)

func DefineFlags() {
	flag.Bool(Server, false, "start the server")
	flag.Uint(Port, DefaultPort, "port to listen on")
	flag.String(User, DefaultUser, "user to access")
	flag.String(Secret, DefaultSecret, "secret")
}
