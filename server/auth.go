package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/Yoseph-code/haken/config"
)

func (s *Server) authenticate(c net.Conn) error {
	reader := bufio.NewReader(c)
	authLine, err := reader.ReadString('\n')

	if err != nil {
		return err
	}

	credentials := strings.TrimSpace(authLine)

	parts := strings.Split(credentials, ":")

	if len(parts) != 2 {
		return nil
	}

	username, password := parts[0], parts[1]

	if username != config.DefaultUser || password != config.DefaultSecret {
		return fmt.Errorf("invalid credentials")
	}

	return nil
}
