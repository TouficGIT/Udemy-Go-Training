package main

import (
	"fmt"
	"io/ioutil"
)

func readFile(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	// On retourne le plus vite possible l'erreur, sil y en a 1
	// -> c'est le "Early return !"
	if err != nil {
		return "", err
	}

	if len(data) == 0 {
		//return "", errors.New("Empty content")
		return "", fmt.Errorf("Empty content (filename=%v)", filename)
	}

	return string(data), err
}

func main() {
	data, err := readFile("test.txt")
	if err != nil {
		fmt.Printf("Error while reading file: %v\n", err)
		return
	}
	fmt.Println("File content:")
	fmt.Println(data)
}
