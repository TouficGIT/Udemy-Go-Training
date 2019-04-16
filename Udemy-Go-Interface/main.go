/*

** Définition en go:
Interface = regroupement nommé de signatures de fonctions

** Régles et bonnes pratiques :
- Une interface ne contient que des signatures de fonctions
- faire des petites interfaces (découpage en pls interface)
- Nommage :
	- Si une interface à 1 seule fonction: Suffixe + er à la fin
	(exemple: type Saver interface { Save(u User) error })
	- Composer des interfaces si possible
	(exemple: type UserSaverLoader interface {})
*/

package main

import "fmt"

// Définition d'un type Instrumenter avec la fonction Play

// ********** Partie I ********** //
type Instrumenter interface {
	Play()
}
type Amplifier interface {
	Connect(amp string)
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
func (g Guitar) Connect(amp string) {
	fmt.Printf("Connected to %v", amp)
}

// ********** Partie II ********** //
type Animal interface {
	Speak() string
}

type Dog struct {
	color string
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
	jumpHeight int
}

func (c Cat) Speak() string {
	return "Meow!"
}

func describeAnimal(a Animal) {
	switch v := a.(type) {
	case Dog:
		fmt.Printf("We have a dog, color = %v\n", v.color)
	case Cat:
		fmt.Printf("We have a cat, jumpHeight = %v\n", v.jumpHeight)
	}
}

func main() {

	// ********** Partie I : intéret d'une interface ********** //
	var instr Instrumenter
	// on instancie un guitar avec 5 cordes
	instr = &Guitar{5}
	instr.Play()
	// instr.Connect() n'est pas possible
	// car un instrument ne connait de fonction connect

	g := Guitar{6}
	g.Play()
	g.Connect("Marshall\n")

	// ********** Partie II : comportement en fonction du type ********** //
	animals := []Animal{
		Dog{"white"},
		Cat{2},
	}
	for _, a := range animals {
		fmt.Println(a.Speak())
		fmt.Printf("Type of animal ? %T\n", a)
		// Type assertion

		//t, ok := a.(Dog)
		//fmt.Printf("Type assertion value=%v, ok=%v\n", t, ok)
		if t, ok := a.(Dog); ok {
			fmt.Printf("We have a dog! color = %v\n", t.color)
		} else {
			fmt.Println("It's not a dog here ...")
		}
	}
	fmt.Println("-------------------------")
	for _, a := range animals {
		describeAnimal(a)
	}

	// ********** Partie III : empty interface ********** //
}
