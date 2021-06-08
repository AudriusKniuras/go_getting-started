package main

import (
	"fmt"

	"github.com/AudriusKniuras/go_getting-started/greetings"
)

func main() {
	message := greetings.Hello("Audrius")
	fmt.Println(message)
}
