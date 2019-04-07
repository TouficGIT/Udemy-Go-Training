package main

import (
	"fmt"
	"strings"
)

// Value receiver
type Rect struct {
	Width, Height int
}

func (r Rect) Area() int {
	return r.Width * r.Height
}

// !: impossible to change a struct in a receiver
// the result is a copy
func (r Rect) Double() {
	r.Width *= 2
	r.Height *= 2
	fmt.Printf("Double() %v\n", r)
}

// Shaping of a struct
func (r Rect) String() string {
	return fmt.Sprintf("Rect ==> widht= %v, height= %v", r.Width, r.Height)
}

// Pointers: exercise function
func UpdateVal(val string) {
	val = "value"
}

func UpdatePtr(ptr *string) {
	*ptr = "pointer"
}

func main() {

	// ************************************************************ //
	// Les "value receiver" en Go
	// ************************************************************ //
	fmt.Println("")
	fmt.Println(strings.ToUpper("Les 'value receiver' en Go"))
	fmt.Println("")
	r := Rect{2, 4}
	fmt.Printf("Rect area= %v \n", r.Area())
	fmt.Println(r)

	// !: Il n'est pas possible de modifer le struct dans la méthode
	r.Double()
	fmt.Println("Main() ", r)

	// ************************************************************ //
	// Les pointeurs en Go
	// ************************************************************ //
	fmt.Println("")
	fmt.Println(strings.ToUpper("Les pointeurs en Go"))
	fmt.Println("")
	i := 1
	// p contient maintenant l'@ de i
	var p *int = &i
	fmt.Printf("i= %v\n", i)
	fmt.Printf("p= %v\n", p)
	fmt.Printf("*p= %v\n", *p)
	fmt.Println("-------------")
	fmt.Println("")
	s := "Bob"
	sPtr := &s
	s2 := *sPtr
	fmt.Println("String pointer")
	fmt.Printf("s= %v\n", s)
	fmt.Printf("sPtr= %v\n", sPtr)
	fmt.Printf("*sPtr= %v\n", *sPtr)
	fmt.Printf("s2= %v\n", s2)
	fmt.Println("-------------")
	fmt.Println("")

	*sPtr = "Alice"
	fmt.Println("Dereference and update")
	fmt.Printf("s= %v\n", s)
	fmt.Printf("sPtr= %v\n", sPtr)
	fmt.Printf("*sPtr= %v\n", *sPtr)
	fmt.Printf("s2= %v\n", s2)
	fmt.Println("-------------")
	fmt.Println("")

	UpdateVal(s)
	fmt.Println("Func UpdateVal")
	fmt.Printf("s= %v\n", s)
	fmt.Printf("*sPtr= %v\n", *sPtr)
	fmt.Println("-------------")
	UpdatePtr(&s)
	fmt.Println("Func UpdatePtr")
	fmt.Printf("s= %v\n", s)
	fmt.Printf("*sPtr= %v\n", *sPtr)
	fmt.Println("-------------")

	// ************************************************************ //
	// La combinaison de pointer et de receiver
	// ************************************************************ //
	fmt.Println("")
	fmt.Println(strings.ToUpper("La combinaison de pointer et de receiver"))
	fmt.Println("")
}
