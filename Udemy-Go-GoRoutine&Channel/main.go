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
package main

import (
	"fmt"
	"time"
)

func longOperation(duration int) {
	time.Sleep(time.Duration(duration) * time.Second)
	fmt.Printf("Operation finished! Took %d s\n", duration)
}

func main() {
	fmt.Println("Starting first operation")
	go longOperation(2)
	fmt.Println("Starting second operation")
	go longOperation(1)
	time.Sleep(2 * time.Second)
}
