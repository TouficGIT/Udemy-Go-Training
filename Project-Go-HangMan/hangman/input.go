package hangman

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Création d'un nouveau reader => recup lettre depuis l'entrée std
var reader = bufio.NewReader(os.Stdin)

// ReadGuess : permet de lire l'entrée de l'utilisateur
// cela retourne => la supposition du joueur (guess)
func ReadGuess() (guess string, err error) {
	valid := false
	for !valid {
		fmt.Print("What is your letter ? ")
		// On passe 'n' dans readString => on lit l'entrée
		// une fois que l'utilisateur appui sur entrée.
		guess, err = reader.ReadString('\n')
		if err != nil {
			return guess, err
		}
		guess = strings.TrimSpace(guess)
		if len(guess) != 1 {
			fmt.Printf("Invalid letter input. letter:%v\n", guess)
			continue
		}
		valid = true
	}
	return guess, nil
}
