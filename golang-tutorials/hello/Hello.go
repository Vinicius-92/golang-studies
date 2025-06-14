package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// single hello
	message, err := greetings.Hello("Vinicius")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	// multiple hello's
	names := []string{"Vinicius", "Johnny", "Sushi", "Freddy"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)

}
