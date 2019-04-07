package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FindReplaceFile(src, dst, old, new string) (occ int, lines []int, err error) {

	occInit := 0
	numLine := 0

	srcFile, err := os.Open(src)
	if err != nil {
		return 0, nil, fmt.Errorf("Impossible to open: %v", srcFile)
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return 0, nil, fmt.Errorf("Impossible to create: %v", dstFile)
	}
	defer dstFile.Close()

	scanner := bufio.NewScanner(srcFile)
	for scanner.Scan() {
		numLine++
		line := scanner.Text() // line contient 1 ligne
		found, res, occ := ProcessLine(line, old, new)
		if found != false {
			lines = append(lines, numLine)
			occInit += occ
			fmt.Printf("Line %d : %v\n", numLine, res)
		} else {
			fmt.Printf("Line %d : %v\n", numLine, res)
		}
		writer := bufio.NewWriter(dstFile)
		defer writer.Flush()
		fmt.Fprintln(writer, res)
	}
	return occInit, lines, err
}

func ProcessLine(line, old, new string) (found bool, res string, occ int) {
	found = strings.Contains(line, old)
	if found != false {
		occ = strings.Count(line, old)
		for i := 0; i < occ; i++ {
			res = strings.Replace(line, old, new, occ)
		}
	} else {
		res = line
		occ = 0
	}
	return found, res, occ
}

func main() {
	occ, lines, err := FindReplaceFile("test.txt", "dst.txt", "Go ", "Python ")
	if err != nil {
		return
	}
	fmt.Println("")
	fmt.Printf("Nb occ: %v\n", occ)
	fmt.Printf("In lines: %v\n", lines)
}
