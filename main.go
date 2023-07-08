package main

import (
	"fmt"
	"log"

	"github.com/Myertek/go-test/db"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	message, err := db.Hello("")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
