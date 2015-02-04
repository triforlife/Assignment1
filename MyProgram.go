// MyProgram.go
package main

import "fmt"

// The struct
type Countdown struct {
	race  string
	go1   string
	count string
	final string
}

// The main function
func main() {
	message := "\nThis is a race program..."
	var address *string = &message
	fmt.Println(message, "This message brought to you from: ", address) // address is pointing to the memory location

	var s = Countdown{"\nLet the race begin...", "Go Go Go\n", "3..2..1..", "At the finish line... You're the Winner!!!"}
	fmt.Println(s.race)
	fmt.Println(s.count, s.go1)
	fmt.Println(s.final)
}
