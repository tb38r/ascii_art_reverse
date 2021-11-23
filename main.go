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

	// test := []string{}

	// for i := 0; i < 8; i++ {
	// 	test = append(test, (lines[i][:len(asciiMap[35][0])]))
	// }
	// fmt.Println(test)

	// for i := 0; i < 8; i++ {
	// 	fmt.Println(lines[i][:len(asciiMap[35][0])])
	// }

	for i := len(lines[0]); i >= 0; i-- {
		for key, value := range asciiMap {
			for ln, line := range value {
				for _, ch := range line {
					fmt.Println(ln, ch, key)
				}
			}
		}
	}

	// // loop over given ascii slice
	// loop over map, char by char, compare the runes of each char to the same part in given
	// test2 := []string{}
	// a := []string{}

	// for i, value := range asciiMap {
	// 	if Compare(value, test) {
	// 		fmt.Println(string(rune(i)))
	// 	}
	// }
	// fmt.Println(a)

	// a := []string{}
	// fmt.Println(len(asciiMap[33][0]))
	// for i := 0; i < 8; i++ {
	// 	a = append(a, lines[i][len(asciiMap[35][0]):]+"\n")
	// }
	// fmt.Println(a)
	// // fmt.Println(lines[0][:len(asciiMap[35][0])])
	// fmt.Println(len(asciiMap[35][0]))

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

// putting the runes of each char from standard into a slice
// runes := []rune{}
// for _, line := range asciiMap[33] {
// 	for _, ch := range line {
// 		runes = append(runes, ch)
// 	}
// }

// runes2 := []rune{}

// // putting given ascii into a slice of runes
// for _, line := range lines {
// 	for _, ch := range line {
// 		runes2 = append(runes2, ch)
// 	}
// }
