package hangman

import "fmt"

func DrawWelcome() {
	fmt.Println(` 
	              _______  _        _______  _______  _______  _       
		|\     /|(  ___  )( (    /|(  ____ \(       )(  ___  )( (    /|
		| )   ( || (   ) ||  \  ( || (    \/| () () || (   ) ||  \  ( |
		| (___) || (___) ||   \ | || |      | || || || (___) ||   \ | |
		|  ___  ||  ___  || (\ \) || | ____ | |(_)| ||  ___  || (\ \) |
		| (   ) || (   ) || | \   || | \_  )| |   | || (   ) || | \   |
		| )   ( || )   ( || )  \  || (___) || )   ( || )   ( || )  \  |
		|/     \||/     \||/    )_)(_______)|/     \||/     \||/    )_)

	`)

}

func Draw(g *Game, guess string) {
	drawTurns(g.TurnsLeft)
	drawState(g, guess)
}

func drawTurns(l int) {
	var draw string
	switch l {
	case 0:
		draw = `
	 +---+
	 |   |
	 O   |
	/|\  |
	/ \  |
	   	 |
	=========
	  `
	case 1:
		draw = `
	 +---+
	 |   |
	 O   |
	/|\  |
	     |
	   	 |
	=========
	  `
	case 2:
		draw = `
	 +---+
	 |   |
	 O   |
	     |
	     |
	   	 |
	=========
	  `
	case 3:
		draw = `
	 +---+
	 |   |
	     |
	     |
	     |
	   	 |
	=========
	  `
	case 4:
		draw = `
	 +---+
	     |
	     |
	     |
	     |
	   	 |
	=========
	  `
	case 5:
		draw = `
	
	     |
	     |
	     |
	     |
	   	 |
	=========
	  `
	case 6:
		draw = `

	     |
	     |
	   	 |
	=========
	  `
	case 7:
		draw = `
	=========
	  `
	case 8:
		draw = `

	  `
	}
	fmt.Println(draw)
}

func drawState(g *Game, guess string) {
	fmt.Print("Guessed: ")

}
