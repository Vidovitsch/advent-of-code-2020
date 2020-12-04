package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readPuzzleInput(path string) []string {
	puzzleInput := []string{}

	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		puzzleInput = append(puzzleInput, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return puzzleInput
}

func solve(puzzleInput []string, horInc int, verInc int) int {
	count := 0
	row, col := 1, 1

	for row <= len(puzzleInput) {
		if string(puzzleInput[row-1][col-1]) == "#" {
			count++
		}
		row += verInc
		col += horInc

		if col > len(puzzleInput[0]) {
			col %= len(puzzleInput[0])
		}
	}

	return count
}

func main() {
	puzzleInput := readPuzzleInput("./../input.txt")

	fmt.Printf("Solution: %v", solve(puzzleInput, 3, 1))
}
