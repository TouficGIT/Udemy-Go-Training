package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// FindReplaceFile
func FindReplaceFile(src, dst, old, new string) (occ int, lines []int, err error) {
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

	writer := bufio.NewWriter(dstFile)
	defer writer.Flush()

	occInit := 0
	numLine := 0

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
		fmt.Fprintln(writer, res)
	}
	return occInit, lines, err
}

// ProcessLine searches for old in line to replace it by new
// It returns found=true, if the pattern was found, res with the resulting string
// and occ with the number of occurence of old
func ProcessLine(line, old, new string) (found bool, res string, occ int) {
	oldLower := strings.ToLower(old)
	newLower := strings.ToLower(new)
	res = line
	if strings.Contains(line, old) || strings.Contains(line, oldLower) {
		found = true
		occ += strings.Count(line, old)
		occ += strings.Count(line, oldLower)
		res = strings.Replace(line, old, new, -1)
		res = strings.Replace(res, oldLower, newLower, -1)
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
