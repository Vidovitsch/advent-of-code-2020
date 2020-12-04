package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type passport struct {
	value map[string]string
}

func (p *passport) isValid() bool {
	for _, requiredField := range []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} {
		_, exists := p.value[requiredField]
		if !exists {
			return false
		}
	}

	return true
}

func newPassport(passportString string) *passport {
	passportStructured := map[string]string{}
	for _, pair := range strings.Split(strings.ReplaceAll(passportString, "\n", " "), " ") {
		splitPair := strings.Split(pair, ":")
		passportStructured[splitPair[0]] = splitPair[1]
	}

	return &passport{value: passportStructured}
}

func readPuzzleInput(path string) []string {
	input, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return strings.Split(string(input), "\n\n")
}

func solve(puzzleInput []string) int {
	valid := 0
	for _, passportString := range puzzleInput {
		if newPassport(passportString).isValid() {
			valid++
		}
	}
	return valid
}

func main() {
	puzzleInput := readPuzzleInput("./../input.txt")

	solution := solve(puzzleInput)

	fmt.Printf("Solution: %v", solution)
}
