package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FindReplaceFile(src, old, new string) (occ int, lines []int, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return 0, nil, fmt.Errorf("Impossible to open: %v", srcFile)
	}
	defer srcFile.Close()

	scanner := bufio.NewScanner(srcFile)
	for scanner.Scan() {
		line := scanner.Text() // line contient 1 ligne
		found, res, occ := ProcessLine(line, old, new)
	}
	return occ, lines, err
}

func ProcessLine(line, old, new string) (found bool, res string, occ int) {
	found = strings.Contains(line, old)
	if found == true {
		occ = strings.Count(line, old)
		res = strings.Replace(line, old, new, 0)
		return
	}
	else {
		res = line
	}
	return found, res, occ
}

func main() {
	FindReplaceFile("test.txt", "Go", "Python")
}
