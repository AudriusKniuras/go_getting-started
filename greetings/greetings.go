package greetings

import "fmt"

// Capital H means exported function (can be called from another package)
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
