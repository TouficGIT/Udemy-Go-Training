package main

import "fmt"

// Définition d'un type Instrumenter avec la fonction Play

type Instrumenter interface {
	Play()
}

// On créé une structure de guitar
type Guitar struct {
	Strings int
}

// On redéfini notre interface en faisant une méthode receiver de Intumenter
// qui va respecter la structure de Play

func (g Guitar) Play() {
	fmt.Printf("Tzzzzoouuuuiiinng with %d strings\n", g.Strings)
}

func main() {
	var instr Instrumenter
	// on instancie un guitar avec 5 cordes
	instr = &Guitar{5}
	instr.Play()
}
