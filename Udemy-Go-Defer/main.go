package main

import "fmt"

func start() {
	fmt.Println("Start")
}

func finish() {
	fmt.Println("Finish")
}

func main() {
	start()
	defer finish() // LIFO = Last In First Out
	// On diffère l'appel de finish au retour de la fonction main()
	// Un defer est rattaché à la fonction (ici main())
	names := []string{"Bob", "Alice", "John"}
	for _, n := range names {
		fmt.Printf("Hello %v\n", n)
	}
}
