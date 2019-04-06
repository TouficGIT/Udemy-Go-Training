package main

import (
	"fmt"
	"strings"

	"Training.go/Udemy-Go-HelloWorld/data"
)

func printInfoNoParam() {
	fmt.Printf("Name = %s , age = %d, email = %s \n", "Bob", 10, "Bob@golang.org")
}

func avg(x, y float64) float64 {
	return (x + y) / 2
}

func titleUpper(title string) {
	fmt.Println(strings.ToUpper(title))
}

func nextChap() {
	fmt.Println("")
}

func toLowerString(name string) (string, int) {
	return strings.ToLower(name), len(name)
}

func main() {

	// ************************************************************ //
	// Print d'une phrase en Go
	// ************************************************************ //
	fmt.Println("Hello world")
	fmt.Println("")

	// ************************************************************ //
	// Utilisation des variables privées ou publique
	// ************************************************************ //
	// fmt.Println(data.password) -> variable privé dans Data, car déclarée avec minuscule
	// fmt.Println(data.Name) -> variable Publique dans Data, car déclarée avec MAJUSCULE
	fmt.Println(strings.ToUpper("Utilisation des variables privées ou publique"))
	fmt.Println(data.Name, data.Age)
	fmt.Println("")

	// ************************************************************ //
	// Utilisation des if et else en Go
	// ************************************************************ //
	fmt.Println(strings.ToUpper("Utilisation des if/else en Go"))
	if data.Age > 10 {
		fmt.Println("Test")
		fmt.Printf("texte formaté, nom: %s\n", data.Name)
		// Printf permet d'afficher des textes formatés
	}

	year, month, day := 2009, 11, 10
	fmt.Printf("Date = %d/%d/%d \n", day, month, year)

	if year == 2009 && month == 11 && day == 10 {
		fmt.Println("This is the first release of Go !")
	} else if year == 2009 || month == 11 || day == 10 {
		fmt.Println("At least one part is right")
	} else {
		fmt.Println("Just another day...")
	}

	if count := 12; count > 10 {
		fmt.Printf("We have enough count, got %d \n", count)
	} // Count uniquement dispo dans ce if
	fmt.Println("")

	// ************************************************************ //
	// Utilisation des switch/cases en Go
	// ************************************************************ //
	fmt.Println(strings.ToUpper("Utilisation des switch/cases en Go"))
	weekDay := 6
	fmt.Printf("Day of the weekd= %d. What's for today? \n", weekDay)

	switch weekDay {
	case 1:
		fmt.Println("Beginning of the week, let's get to work !")
	case 3:
		fmt.Println("Wednesday, the half is done !")
	case 6, 7:
		fmt.Println("It's the weekend !")
	default:
		fmt.Println("Nothing special")
	}

	hour := 10
	fmt.Printf("Current time= %d\n", hour)
	switch {
	case hour < 12:
		fmt.Println("It's the morning")
	case hour > 12 && hour < 18:
		fmt.Println("It's the afternoon")
	default:
		fmt.Println("It's the evening")
	}
	fmt.Println("")

	// ************************************************************ //
	// Conversion de types en Go
	// ************************************************************ //
	fmt.Println(strings.ToUpper("Conversion de types en GO"))
	// Conversion float en int
	var percentage float32 = 2.0
	fmt.Printf("Current percentage %f%%\n", percentage)
	fmt.Printf("Int percentage = %d%%\n", int(percentage))

	// Conversion int en float
	n := 42
	fmt.Printf("int= %d\n", n)
	f := float64(n) + 0.42
	fmt.Printf("float(int+0.42)= %f\n", f)

	fmt.Println("")

	// ************************************************************ //
	// Les fonctions en Go
	// ************************************************************ //
	titleUpper("Les fonctions en Go")
	fmt.Println("Test de ma nouvelle fonction de titre")
	printInfoNoParam()
	result := avg(16.3, 30.5)
	fmt.Printf("La moyenne : %f \n", result)
	nextChap()

	// ************************************************************ //
	// Les retours de fonction multiples en Go
	// ************************************************************ //
	titleUpper("Les retours de fonction multiples en Go")
	lowerName, lenght := toLowerString("ALICE")
	fmt.Printf("Nom lower case : %s et taille : %d \n", lowerName, lenght)
	_, lenght = toLowerString("Bob") // On peut choisir de ne prendre qu'un seul param de retour
	fmt.Printf("On affiche la taille uniquement: %d \n", lenght)
	nextChap()

	// ************************************************************ //
	// Les tableaux en Go
	// ************************************************************ //
	titleUpper("Les tableaux en Go")
	var names [3]string
	fmt.Printf("names= %v (len = %v) \n", names, len(names))
	names[0] = "Bob"
	names[2] = "Alice"
	fmt.Printf("names= %v (len = %v) \n", names, len(names))
	fmt.Printf("names[2]= %v \n", names[2])
	odds := [4]int{1, 3, 5, 7} // déclaration + remplissage direct d'1 tab
	fmt.Printf("odds=%v (len=%d)\n", odds, len(odds))
	nextChap()

	// ************************************************************ //
	// Les slices en Go
	// ************************************************************ //
	titleUpper("Les slices en Go")
	// Slice ~= array de taille dynamique
	nums := make([]int, 2, 3)
	nums[0] = 10
	nums[1] = -2
	fmt.Printf("nums=%v (len=%d) (cap=%d)\n", nums, len(nums), cap(nums))
	nums = append(nums, 64)
	fmt.Printf("nums=%v (len=%d) (cap=%d)\n", nums, len(nums), cap(nums))
	nums = append(nums, -8)
	fmt.Printf("nums=%v (len=%d) (cap=%d)\n", nums, len(nums), cap(nums))

	// Test d'ajout ds 1 slice -> modification de la taille et de la cap
	letters := []string{"g", "o", "l", "a", "n", "g"}
	fmt.Printf("letters=%v (len=%d) (cap=%d)\n", letters, len(letters), cap(letters))
	letters = append(letters, "!")
	fmt.Printf("letters=%v (len=%d) (cap=%d)\n", letters, len(letters), cap(letters))

	// Test d'une subdivision d'un slice
	sub1 := letters[:2]
	sub2 := letters[2:]
	fmt.Printf("sub1=%v (len=%d) (cap=%d)\n", sub1, len(sub1), cap(sub1))
	fmt.Printf("sub2=%v (len=%d) (cap=%d)\n", sub2, len(sub2), cap(sub2))

	sub2[0] = "UP"
	fmt.Printf("letters=%v (len=%d) (cap=%d)\n", letters, len(letters), cap(letters))
	fmt.Printf("sub1=%v (len=%d) (cap=%d)\n", sub1, len(sub1), cap(sub1))
	fmt.Printf("sub2=%v (len=%d) (cap=%d)\n", sub2, len(sub2), cap(sub2))

	// Test d'une copy de slice
	subCopy := make([]string, len(sub1))
	copy(subCopy, sub1)
	subCopy[0] = "COPY"
	fmt.Printf("subCopy=%v (len=%d) (cap=%d)\n", subCopy, len(subCopy), cap(subCopy))
	nextChap()

	// ************************************************************ //
	// Les boucles for en Go
	// ************************************************************ //
	titleUpper("Les boucles for en Go")
	// Syntaxe classique
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += 5
	}
	fmt.Printf("i = %v\n", sum)

	// Syntaxe While
	eventCount := 0
	for eventCount < 3 {
		fmt.Println("Retrieving events...")
		eventCount++
		if eventCount == 3 {
			fmt.Printf("Got %d notifications \n", eventCount)
		}
	}
	// Syntaxe Forever
	i := 0
	for {
		i++
		if i%2 != 0 {
			fmt.Println("Odd looping...")
			continue
		}
		fmt.Println("Looping...")
		if i >= 4 {
			fmt.Println("Loop end")
			break
		}
	}
	nextChap()

	// ************************************************************ //
	// Les ranges en Go
	// ************************************************************ //
	titleUpper("Les ranges en Go")
	// Permet d'itérer sur les éléments d'1 struct de données
	userNames := []string{"Bob", "Alice", "John"}
	for i, n := range userNames {
		fmt.Printf("Usernames=%s (index= %d)\n", n, i)
	}
	for _, c := range "golang" {
		fmt.Printf("%v\n", string(c))
	}
	nextChap()
}
