package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// opening the given ascii file in read only mode. The file must exist(in the current working directory)
	// need os.Args[1][10:] later for file name
	givenfile, _ := os.Open("given.txt")
	scanned := bufio.NewScanner(givenfile)

	scanned.Split(bufio.ScanLines)

	var lines []string

	// append each line into a slice of strings
	for scanned.Scan() {
		lines = append(lines, scanned.Text())
	}
	givenfile.Close()

	// opening the standard txt file with all the chars, appending all lines to a slice of string
	asciifile, _ := os.Open("standard.txt")
	scanLines := bufio.NewScanner(asciifile)

	scanLines.Split(bufio.ScanLines)

	var asciiLines []string

	for scanLines.Scan() {
		asciiLines = append(asciiLines, scanLines.Text())
	}

	asciifile.Close()

	// create a map and append all ascii chars, line number corresponding to ascii number
	asciiMap := make(map[int][]string)

	asciiNum := 31

	for _, line := range asciiLines {
		if line == "" {
			asciiNum++
		} else {
			asciiMap[asciiNum] = append(asciiMap[asciiNum], line)
		}
	}

	// putting the runes of each char from standard into a slice
	runes := []rune{}
	for _, line := range asciiMap[33] {
		for _, ch := range line {
			runes = append(runes, ch)
		}
	}

	runes2 := []rune{}

	// putting given ascii into a slice of runes
	for _, line := range lines {
		for _, ch := range line {
			runes2 = append(runes2, ch)
		}
	}

	// loop over given ascii slice
	// loop over map, char by char, compare the runes of each char to the same part in given
	for i, value := range asciiMap {
		for j := 0; j < 8; j++ {
			if Compare(value, []string{"hello"}) {
				fmt.Println(string(rune(i)))
				fmt.Println(value, lines[:len(asciiMap[i])])
			}
		}
	}

	// for i := 0; i < 8; i++ {
	// 	fmt.Print(lines[i][:len(asciiMap[35][0])])
	// }

	// fmt.Println(lines[0][:len(asciiMap[35][0])])
	// fmt.Println(len(asciiMap[35][0]))
	// fmt.Println("hello")
	// fmt.Println(strings.Join(asciiMap[35], ""))
}

// Compare, compares slices of string to check if they are equal.
func Compare(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
