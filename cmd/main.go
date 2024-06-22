package main

import (
	"flag"
	"log"

	"github.com/Yoseph-code/haken/cli"
	"github.com/Yoseph-code/haken/db"
	"github.com/Yoseph-code/haken/server"
)

const FlagServer string = "server"

const (
	FlagUser string = "u"
	FlagPort string = "p"
)

func init() {
	flag.Uint(FlagPort, server.DefaultListenAddr, "port to listen on")
	flag.String(FlagUser, "", "user to access")
}

func main() {
	isServer := flag.Bool(FlagServer, false, "start the server")

	flag.Parse()

	if *isServer {
		port := flag.Lookup(FlagPort).Value.(flag.Getter).Get().(uint)

		s := server.New(server.Config{
			ListenAddr: port,
		})

		d, err := db.NewDBFile()

		if err != nil {
			log.Panic(err)
		}

		if ok := d.IsDBExists(); ok {
			s.SetDB(d)
		} else {
			file, err := d.CreateDB()

			if err != nil {
				log.Panic(err)
			}

			defer file.Close()

			s.SetDB(d)
		}

		if err := s.Run(); err != nil {
			log.Panic(err)
		}
	} else {
		port := flag.Lookup(FlagPort).Value.(flag.Getter).Get().(uint)

		c := cli.NewCli("admin", "admin", "localhost", port)

		if err := c.Conect(); err != nil {
			log.Panic(err)
		}

		// if err := c.Ping(); err != nil {
		// 	log.Panic(err)
		// }

		if err := c.Run(); err != nil {
			log.Panic(err)
		}

		// reader := bufio.NewReader(os.Stdin)

		// for {
		// 	fmt.Print("haken> ")

		// 	input, err := reader.ReadString('\n')

		// 	if err != nil {
		// 		log.Panic(err)
		// 	}

		// 	input = strings.TrimSpace(input)

		// 	if input == "exit" {
		// 		fmt.Print("Bye!")
		// 		break
		// 	}

		// 	fmt.Println("Você digitou:", input)

		// 	// fmt.Println("Você digitou:", input)

		// 	// if strings.HasPrefix(input, "set") {
		// 	// 	fmt.Println("set")
		// 	// } else if strings.HasPrefix(input, "get") {
		// 	// 	fmt.Println("get")
		// 	// } else if strings.HasPrefix(input, "del") {
		// 	// 	fmt.Println("del")
		// 	// } else {
		// 	// 	fmt.Println("Comando não reconhecido")
		// 	// }

		// 	// fmt.Println("Comando não reconhecido")

		// 	// if strings.HasPrefix(input, "set") {
		// 	// 	fmt.Println("set")
		// }
	}
}

// func main() {
// 	flag.Uint("p", server.DefaultListenAddr, "HTTP network address")

// 	flag.Parse()

// 	addr := flag.Lookup("p").Value.(flag.Getter).Get().(uint)

// 	s := server.New(server.Config{
// 		ListenAddr: addr,
// 	})

// 	fdb, err := db.NewDBFile()

// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	// file, err := fdb.CreateDB()

// 	// if err != nil {
// 	// 	log.Panic(err)
// 	// }

// 	// defer file.Close()

// 	s.SetDB(fdb)

// 	if err := s.Run(); err != nil {
// 		log.Panic(err)
// 	}
// }
