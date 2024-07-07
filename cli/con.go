package cli

import (
	"fmt"
	"net"
)

func (c *Cli) Connect() error {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", c.Host, c.Port))

	if err != nil {
		return fmt.Errorf("error connecting to server: %w", err)
	}

	fmt.Println(c.User, c.Secret)

	fmt.Fprintf(conn, "%s:%s\n", c.User, c.Secret)

	c.con = conn

	return nil
}

func (c *Cli) Ping() error {
	_, err := fmt.Fprintf(c.con, "PING\n")

	if err != nil {
		return fmt.Errorf("error sending ping command to server: %w", err)
	}

	response := make([]byte, 1024)

	n, err := (c.con).Read(response)

	if err != nil {
		return fmt.Errorf("error reading response from server: %w", err)
	}

	if string(response[:n]) != "PONG" {
		return fmt.Errorf("unexpected response from server: %s", string(response[:n]))
	}

	return nil
}

func (c *Cli) SendCommand(cmd string) (string, error) {
	_, err := fmt.Fprintf(c.con, "%s\n", cmd)

	if err != nil {
		return "", fmt.Errorf("error sending command to server: %w", err)
	}

	response := make([]byte, 1024)

	n, err := (c.con).Read(response)

	if err != nil {
		return "", fmt.Errorf("error reading response from server: %w", err)
	}

	return string(response[:n]), nil
}

func (c *Cli) Close() error {
	err := c.con.Close()

	if err != nil {
		return fmt.Errorf("error closing connection: %w", err)
	}

	return nil
}

// func (c *Cli) SendCommand(cmd string) error {
// 	_, err := fmt.Fprintf(*c.con, "%s\n", cmd)

// 	if err != nil {
// 		return fmt.Errorf("error sending command to server: %w", err)
// 	}

// 	return nil
// }
