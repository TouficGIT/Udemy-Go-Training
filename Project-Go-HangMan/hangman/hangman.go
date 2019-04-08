package hangman

import "strings"

// Game : structure de notre jeu
type Game struct {
	State        string   // Etat du jeu: perdu ou gagné
	Letters      []string // Les lettres a trouver dans le mot
	FoundLetters []string // Bonnes propositions de lettres
	UsedLetters  []string // Les lettres utilisés
	TurnsLeft    int      // Nombres d'essais restants
}

func New(turns int, word string) *Game {
	letters := strings.Split(strings.ToUpper(word), "") // Split: permet de séparer lettres - renvoi un slice
	found := make([]string, len(letters))
	for i := 0; i < len(letters); i++ {
		found[i] = "_"
	}
	g := &Game{
		State:        "",
		Letters:      letters,
		FoundLetters: found,
		UsedLetters:  []string{},
		TurnsLeft:    turns,
	}
	return g
}
