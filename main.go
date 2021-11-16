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

	// for _, line := range lines {
	// 	for _, ch := range line {
	// 		fmt.Println(rune(ch))
	// 	}
	// }

	// figure a way to print the chars vertically

	for j := 0; j < 81; j++ {
		for i := 0; i < len(lines); i++ {
			if lines[i][j] != 32 {
				fmt.Print(rune(lines[i][j]))
				// fmt.Print("\n")
			}
		}
		fmt.Println()
	}

	// for _, line := range asciiMap[72] {
	// 	for _, ch := range line {
	// 		if ch != 32 {
	// 			fmt.Print(rune(ch))
	// 			fmt.Print("\n")
	// 		}
	// 	}
	// }

	// for i := 0; i < 8; i++ {
	// 	for _, ch := range asciiMap[72][i] {
	// 		if ch != 32 {
	// 			fmt.Println(rune(ch))
	// 		}
	// 	}
	// }

	// for i := 0; i < 8; i++ {
	// 	for j, line := range asciiLines {
	// 		if line != "" && asciiLines[j][i] != 32 {
	// 			fmt.Print(asciiLines[j][i])
	// 		}
	// 	}
	// }
}
