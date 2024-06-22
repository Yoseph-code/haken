package cli

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

type Cli struct {
	User   string
	Secret string
	Host   string
	Port   uint

	con net.Conn
}

func NewCli(u, s, h string, port uint) *Cli {
	return &Cli{
		User:   u,
		Secret: s,
		Host:   h,
		Port:   port,
	}
}

func (c *Cli) Run() error {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("haken> ")

		text, err := reader.ReadString('\n')

		if err != nil {
			return fmt.Errorf("error reading input: %w", err)
		}

		text = strings.TrimSpace(text)

		if text == "exit" {
			fmt.Println("Bye!")

			err := c.con.Close()

			if err != nil {
				return err
			}

			break
		}

		command := strings.Fields(text)

		if len(command) == 0 {
			continue
		}

		token := NewToken(strings.ToUpper(command[0]))

		if token.IsInvalid() {
			fmt.Println("Invalid command. Type 'help' to see the list of commands.")
			continue
		}

		switch token {
		case Read:
			if len(command) < 2 {
				fmt.Println("Usage: READ <key>")
				continue
			}
			key := command[1]

			res, err := c.SendCommand(fmt.Sprintf("READ %s", key))

			if err != nil {
				fmt.Println("Error sending command to server:", err)
			}

			if strings.Contains(res, "OK") {
				res = strings.Split(res, "\n")[0]
			} else {
				fmt.Println("Invalid response from server")
				continue
			}

			c.Print(res)
		case Create:
			if len(command) < 3 {
				fmt.Println("Usage: CREATE <key> <value>")
				continue
			}

			key := command[1]

			value := strings.Join(command[2:], " ")

			res, err := c.SendCommand(fmt.Sprintf("CREATE %s %s", key, value))

			if err != nil {
				fmt.Println("Error sending command to server:", err)
			}

			if strings.Contains(res, "OK") {
				res = strings.Split(res, "\n")[0]
			} else {
				fmt.Println("Invalid response from server")
				continue
			}

			c.Print(res)
		case Update:
			fmt.Println("UPDATE command executed")
		case Remove:
			fmt.Println("REMOVE command executed")
		case Help:
			printHelp()
		case Ping:
			res, err := c.SendCommand("PING")

			if err != nil {
				fmt.Println("Error sending command to server:", err)
			}

			if strings.Contains(res, "PONG") {
				res = strings.Split(res, "\n")[0]
			} else {
				fmt.Println("Invalid response from server")
				continue
			}

			c.Print(res)
		default:
			fmt.Println("Unknown command. Type 'help' to see the list of commands.")
		}
	}

	return nil
}

func (c *Cli) Print(res string) {
	showRes := fmt.Sprintf("|        %s        |", res)

	fmt.Println(strings.Repeat("-", len(showRes)))
	fmt.Println(showRes)
	fmt.Println(strings.Repeat("-", len(showRes)))
}

// Token representa os diferentes comandos suportados pela CL

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// type Cli struct {
// 	User   string
// 	Secret string
// 	Host   string
// 	Port   uint
// }

// func NewCli(u, s, h string, port uint) *Cli {
// 	return &Cli{
// 		User:   u,
// 		Secret: s,
// 		Host:   h,
// 		Port:   port,
// 	}
// }

// func (c *Cli) Run() error {
// 	reader := bufio.NewReader(os.Stdin)

// 	for {
// 		fmt.Print("haken>")

// 		text, err := reader.ReadString('\n')

// 		if err != nil {
// 			return err
// 		}

// 		text = strings.TrimSpace(text)

// 		if text == "exit" {
// 			fmt.Println("Bye!")
// 			break
// 		}

// 		command := strings.Fields(text)

// 		if len(command) == 0 {
// 			continue
// 		}

// 		token := NewToken(strings.ToUpper(command[0]))

// 		if token.IsInvalid() {
// 			fmt.Println("Invalid command")
// 			continue
// 		}

// 		switch token {
// 		case Read:
// 			key := ""
// 			if len(command) > 1 {
// 				key = command[1]
// 			} else {
// 				fmt.Println("Invalid command")
// 				continue
// 			}
// 			fmt.Println("READ command executed with key:", key)
// 		case Create:
// 			fmt.Println("CREATE command executed")
// 		case Update:
// 			fmt.Println("UPDATE command executed")
// 		case Remove:
// 			fmt.Println("REMOVE command executed")
// 		}
// 	}

// 	return nil
// }
