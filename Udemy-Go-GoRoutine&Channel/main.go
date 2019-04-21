/*
**** Définition Goroutines : ****

Concurrence: alternance d'éxecution entre chaque programme, les entités en concurrence essaient d'accèder
à des données partagés (risques de problèmes)

Ds la plupart des langages, la concurrence est faites avec les threads (fil d'exécution)
1 process = pls threads

** Race condition : 2 threads essayent de modifier 1 variable en meme tps **
** Starvation : 1 thread ne relâche jamais son accès à une variable **

**** Syntaxe : ****

<prefix "go"> + <function call>
go LongOperation(1)

**** Goroutines ? : ****
Une goroutine est géré par le moteur go, elle ne s'appuie pas sur le mécanisme des threads de l'OS !
Conséquences :
 - On peut éxécuter plus de Goroutines que de threads
 - Les goroutines démarrent + rapidement
 - Empreinte mémoire + faible
 - Pas de changement de contexte que l'OS doit exécuter pour la création

**** Limitation : ****
 - Capacités réduites, pas de garanti de démarrage
 - Sécurité "+ faible"
*/

/*
**** Définition Channels : ****
Comment synchroniser les goroutines ?

Ce sont des tuyaux, qui font communiquer les goroutines.
La lecture / ecriture est synchrone

**** Syntaxe : ****
c := make(chan <Type>)
exemple :
c := make(chan int)
c <- 1 // écrire dans un channel
i := <-c // lire ds un channel
*/

/**** Définition Deadlock : ****
Situation de plantage de notre appli lorsqu'une Goroutine interagit avec un channel
mais qu'il n'y personne pr recevoir ou emettre
*/

/**** Définition Channel Unidirectionnel : ****
Info supplémentaire sur un channel + sécuriser son utilisation
Choisir si on peut lire ou écrire
*/

/**** Le Select : ****
Mécanisme très utile avec les Goroutine et les channels.
C'est une sorte de switch spécial channels.
Il va bloquer jusqu'à ce qu'une lecture/écriture soit prête
*/
package main

import (
	"fmt"
	"time"
)

func longOperation(duration int) {
	time.Sleep(time.Duration(duration) * time.Second)
	fmt.Printf("Operation finished! Took %d s\n", duration)
}

func ping(c chan string) {
	for i := 1; ; i++ {
		c <- fmt.Sprintf("ping %v", i)
	}
}

func pong(c chan string) {
	for i := 100; ; i = i + 100 {
		c <- fmt.Sprintf("pong %v", i)
	}
}

func print(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(1 * time.Second)
	}
}

func hello(c chan string) {
	c <- "Hello there!"
	fmt.Println("hello() finished")
	fmt.Println(<-c)
}

func client1(c chan string) {
	for i := 0; i < 5; i++ {
		c <- fmt.Sprintf("Message from client1 => %v", i)
		time.Sleep(1500 * time.Millisecond)
	}
}

func client2(c chan string) {
	for i := 10; i < 15; i++ {
		c <- fmt.Sprintf("Message from client2 => %v", i)
		time.Sleep(3000 * time.Millisecond)
	}
}

func main() {
	/* Goroutines :
	fmt.Println("Starting first operation")
	go longOperation(2)
	fmt.Println("Starting second operation")
	go longOperation(1)
	time.Sleep(2 * time.Second)
	*/

	/* Channels :
	c := make(chan string)
	go ping(c)
	go pong(c)
	go print(c)

	time.Sleep(10 * time.Second)
	*/

	/* Deadlocks:
	c := make(chan string)
	go hello(c)
	fmt.Println(<-c)
	c <- "Msg from main"
	time.Sleep(1 * time.Second)
	*/

	/* Selects: */
	c1 := make(chan string)
	c2 := make(chan string)
	go client1(c1)
	go client2(c2)

	maxEmpty := 10
	currEmpty := 0
	for currEmpty <= maxEmpty {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-c1:
			fmt.Println("Received from client1: ", v)
		case v := <-c2:
			fmt.Println("Received from client2: ", v)
		default:
			fmt.Println("No value received")
			currEmpty++
		}
	}
}
