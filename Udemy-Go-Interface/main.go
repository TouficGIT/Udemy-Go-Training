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

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// ********** Partie I : intéret d'une interface ********** //

// Définition d'un type Instrumenter avec la fonction Play
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

// ********** Partie II : comportement en fonction du type ********** //
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

// ********** Partie III : empty interface ********** //

type Person struct {
	Name string
	Age  int
}

type Cooker interface {
	Cook()
}

type Chef struct {
	Person
	Stars int
}

func (c Chef) Cook() {
	fmt.Printf("I'm cooking with %v stars\n", c.Stars)
}

func processPerson(i interface{}) {
	switch p := i.(type) {
	case Person:
		fmt.Printf("We have a person= %v\n", p)
	case Chef:
		fmt.Printf("We have a chef=%v, Let him cook...\n", p)
		p.Cook()
	default:
		fmt.Printf("Unknown type=%T, value=%v\n", p, p)
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

	// Dans d'autre language il existe des types génériques, pas en go
	// On peut utiliser l'empty interface pour palier ce manque, elle peut contenir tous les types
	// idée : tous les types implémentent au moins 0 méthodes
	// Exemple :
	var x interface{} = Person{"bob", 10}
	fmt.Printf("x type= %T, data=%v\n", x, x)
	s, ok := x.(string)
	fmt.Printf("Person as string ok ? %v. value='%v'\n", ok, s)

	processPerson(x)

	// Rappel : struct embedded ci-dessous
	x = Chef{
		Stars: 4,
		Person: Person{
			Name: "Alice",
			Age:  22,
		},
	}
	processPerson(x)
	processPerson(3)
	processPerson("Emilien")

	// ********** Partie IV : reader & writer ********** //
	// Exemple ci dessous : reader est construit avec une interface vide
	// permettant de lire des données depuis n'importe quelle type de source (fichier ou string ou http...)

	// r := strings.NewReader("Hello Gophers! Readers are awesomes")
	r, err := os.Open("test.txt")
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		fmt.Printf("Error in reader : %v\n", err)
		return
	}
	fmt.Println(string(buf))

	// ********** Partie V : requête HTTP ********** //
	resp, err := http.Get("https://golang.org")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	// Read All appel la fonction read sur resp.Body
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))
	f, _ := os.Create("golang.html")
	defer f.Close()
	io.Copy(f, resp.Body)
}
