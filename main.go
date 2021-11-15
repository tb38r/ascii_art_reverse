package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// opening the file in read only mode. The file must exist(in the current working directory)
	// need os.Args[1][10:] later for file name
	file, _ := os.Open("given.txt")
	scanned := bufio.NewScanner(file)

	// s
	scanned.Split(bufio.ScanLines)

	var lines []string

	// append each line into a slice of strings
	for scanned.Scan() {
		lines = append(lines, scanned.Text())
	}
	file.Close()

	// figure a way to print the chars vertically

	for j := 0; j < 81; j++ {
		for i := 0; i < len(lines); i++ {
			if lines[i][j] != 32 {
				fmt.Print(string(lines[i][j]))
				// fmt.Print("\n")
			}
		}
		fmt.Println()
	}

	// a := strings.Split(lines, "$")
	// fmt.Println(a)
	// fmt.Println((lines[0][0]))
}
