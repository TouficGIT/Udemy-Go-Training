package main

import (
	"fmt"
	"os"

	"Training.go/Project-Go-HangMan/hangman"
	"Training.go/Project-Go-HangMan/hangman/dictionnary"
)

func main() {
	err := dictionnary.Load("words.txt")
	if err != nil {
		fmt.Printf("Could not load dictionnary: %v\n", err)
		os.Exit(1)
	}
	g := hangman.New(8, dictionnary.PickWord())
	hangman.DrawWelcome()
	guess := ""
	for {
		hangman.Draw(g, guess)
		switch g.State {
		case "won", "lost":
			os.Exit(0)
		}

		l, err := hangman.ReadGuess()
		if err != nil {
			fmt.Printf("Could not read from terminal: %v", err)
			os.Exit(1)
		}
		guess = l
		g.MakeAGuess(guess)
	}
}
