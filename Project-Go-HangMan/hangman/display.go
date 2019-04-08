package hangman

import "fmt"

func DrawWelcome() {
	fmt.Println(`     
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
	drawLetters(g.FoundLetters)

	fmt.Print("Used: ")
	drawLetters(g.UsedLetters)

	switch g.State {
	case "goodGuess":
		fmt.Print("Good guess!")
	case "alreadyGuessed":
		fmt.Printf("Letter '%s' was already used", guess)
	case "badGuess":
		fmt.Printf("Bad guess, '%s' is not in the word", guess)
	case "lost":
		fmt.Print("You lost :( the word was:")
		drawLetters(g.Letters)
	case "won":
		fmt.Print("You won! The word was:")
		drawLetters(g.Letters)
	}
}

func drawLetters(l []string) {
	for _, c := range l {
		fmt.Printf("%v: ", c)
	}
	fmt.Println()
}
