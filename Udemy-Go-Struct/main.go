package main

import (
	"fmt"
	"strings"

	"Training.go/Udemy-Go-Struct/Player"
)

// 'Embedded' structs
type User struct {
	Name  string
	Email string
}

// On définit un admin comme un User
// on fait une "composition" ~= à de l'héritage
type Admin struct {
	User
	Level int
}

func main() {

	// ************************************************************ //
	// Les struct "classiques" en Go
	// ************************************************************ //
	fmt.Println("")
	fmt.Println(strings.ToUpper("Les struct 'classiques' en Go"))
	fmt.Println("")
	var p1 Player.Player
	p1.Name = "Bob"
	p1.Age = 10

	fmt.Printf("Player 1: %v\n", p1)
	fmt.Printf("p1.Name: %v, p1.Age= %v\n", p1.Name, p1.Age)

	a := Player.Avatar{"http://avatar.jpg"}
	fmt.Printf("avatar: %v\n", a)

	p3 := Player.Player{
		Name: "John",
		Age:  25,
		Avatar: Player.Avatar{
			Url: "http://url.com",
		},
	}
	fmt.Printf("Player 3: %v\n", p3)

	// Initialisation "correct" d'une struct
	p4 := Player.New("Emilien")
	fmt.Printf("Player 4: %v\n", p4)
	fmt.Println("")

	// ************************************************************ //
	// Les "Embedded" structs en Go
	// ************************************************************ //
	fmt.Println(strings.ToUpper("Les 'Embedded' structs en Go"))
	fmt.Println("")
	u := User{
		Name:  "Emilien",
		Email: "Emilien@golang.org",
	}
	fmt.Printf("User: %v\n", u)

	ad := Admin{
		Level: 2,
		User: User{
			Name:  "Alice",
			Email: "alice@golang.org",
		},
	}
	fmt.Printf("Admin: %v\n", ad)
	// A la déclaration du struct embedded il faut définir la classe "sup"
	// A l'utilisation c'est transparent => cf ci dessous
	fmt.Printf("Admin name: %v, level: %v\n", ad.Name, ad.Level)
	fmt.Println("")
}
