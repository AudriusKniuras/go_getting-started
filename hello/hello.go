package main

import (
	"fmt"
	"log"

	"github.com/AudriusKniuras/go_getting-started/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	messages, err := greetings.Hellos

	message, err := greetings.Hello("Audrius")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
	// slices (lists in python)
	slice_int := []int{1, 2, 3}
	fmt.Println(slice_int[1])
}
