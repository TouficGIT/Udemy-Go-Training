package main

import (
	"flag"
	"fmt"
	"os"

	"Training.go/Project-Go-GenCerts/cert"
	"Training.go/Project-Go-GenCerts/html"
	"Training.go/Project-Go-GenCerts/pdf"
)

func main() {
	// CLI
	outputType := flag.String("type", "pdf", "Output type of the certificate")
	flag.Parse()

	var saver cert.Saver
	var err error
	switch *outputType {
	case "html":
		saver, err = html.New("output")
	case "pdf":
		saver, err = pdf.New("output")
	default:
		fmt.Printf("Unknown output type. got=%v", *outputType)
	}

	if err != nil {
		fmt.Printf("Error during pdf generator creation: %v", err)
		os.Exit(1)
	}
	c, err := cert.New("Golang programming", "Emilien Meffe", "2019-04-17")
	if err != nil {
		fmt.Printf("Error during certificate creation: %v", err)
		os.Exit(1)
	}
	saver.Save(*c)
}
