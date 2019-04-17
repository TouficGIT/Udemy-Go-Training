package main

import (
	"fmt"
	"os"

	"Training.go/Project-Go-GenCerts/cert"
	"Training.go/Project-Go-GenCerts/html"
)

func main() {
	c, err := cert.New("Golang programming", "Emilien Meffe", "2019-04-17")
	if err != nil {
		fmt.Printf("Error during certificate creation: %v", err)
		os.Exit(1)
	}

	var saver cert.Saver
	saver, err = html.New("output")
	if err != nil {
		fmt.Printf("Error during pdf generator creation: %v", err)
		os.Exit(1)
	}
	saver.Save(*c)
}
