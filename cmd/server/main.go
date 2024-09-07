package main

import (
	"flag"
	"fmt"

	"github.com/Yoseph-code/haken/config"
	"github.com/Yoseph-code/haken/hf"
	"github.com/Yoseph-code/haken/internal/users"
)

func init() {
	config.DefineServerFlags()
}

func main() {
	flag.Parse()

	uhf := hf.NewUserHakenFile()

	if err := uhf.DB(); err != nil {
		panic(err)
	}

	u, err := uhf.LoadUsers()

	if err != nil {
		panic(err)
	}

	if len(u) == 0 {
		user := users.NewUser("root", "", users.SuperAdmin)

		if err := user.Valid(); err != nil {
			panic(err)
		}

		if err := uhf.SaveUser(user); err != nil {
			panic(err)
		}
	}

	port := flag.Lookup(config.ServerPort).Value.(flag.Getter).Get().(uint)
	user := flag.Lookup(config.ServerUser).Value.(flag.Getter).Get().(string)
	secret := flag.Lookup(config.ServerSecret).Value.(flag.Getter).Get().(string)
	db := flag.Lookup(config.ServerDB).Value.(flag.Getter).Get().(string)

	fmt.Println(port)
	fmt.Println(user)
	fmt.Println(secret)
	fmt.Println(db)
}
