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
	inputFile := flag.String("file", "", "Input csv file")
	flag.Parse()

	if len(*inputFile) <= 0 {
		fmt.Printf("Invalid file. got='%v'", *inputFile)
		os.Exit(1)
	}

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
	certs, err := cert.ParseCsv(*inputFile)
	if err != nil {
		fmt.Printf("Could not parse CSV file: %v", err)
		os.Exit(1)
	}
	for _, c := range certs {
		err := saver.Save(*c)
		if err != nil {
			fmt.Printf("Could not save Cert. got=%v\n", err)
		}
	}
}
