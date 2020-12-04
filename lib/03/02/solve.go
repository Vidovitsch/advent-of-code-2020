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

	solution := 1
	for _, slope := range [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}} {
		solution *= solve(puzzleInput, slope[0], slope[1])
	}

	fmt.Printf("Solution: %v", solution)
}
