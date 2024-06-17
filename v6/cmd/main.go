package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Yoseph-code/haken/db"
	"github.com/Yoseph-code/haken/three"
)

const (
	t = 3
)

func main() {
	bt := three.NewBThree(t)

	pairs := map[string]string{
		"John":    "Doe",
		"Jane":    "Smith",
		"Michael": "Johnson",
		"Emily":   "Davis",
		"Daniel":  "Wilson",
		"Sophia":  "Garcia",
		"David":   "Martinez",
		"Emma":    "Anderson",
	}

	for k, v := range pairs {
		bt.Insert(k, v)
	}

	bt.Print()

	fmt.Println("--------------------------------")

	name := "John"
	if surname, found := bt.Search(name); found {
		fmt.Printf("Sobrenome de %s é %s\n", name, surname)
	} else {
		fmt.Printf("%s não encontrado\n", name)
	}

	pwd, err := os.Getwd()

	if err != nil {
		log.Panic(err)
	}

	mainPath := filepath.Join(pwd, "haken", "default.bin")

	// err = db.CreateDB(mainPath)

	// if err != nil {
	// 	log.Panic(err)
	// }

	db := db.NewDB(mainPath, t)

	fmt.Println(db)

	err = db.AppendToFile(bt.Root)

	if err != nil {
		log.Panic(err)
	}

	err = db.ReadFromFile()

	if err != nil {
		log.Panic(err)
	}

	db.Data.Print()
}
