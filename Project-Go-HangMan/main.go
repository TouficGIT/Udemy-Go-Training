package main

import (
	"fmt"
	"os"

	"Training.go/Project-Go-HangMan/hangman"
)

func main() {
	g := hangman.New(8, "Golang")
	fmt.Println(g)

	l, err := hangman.ReadGuess()
	if err != nil {
		fmt.Printf("Could not read from terminal: %v", err)
		os.Exit(1)
	}
	fmt.Println(l)
}
