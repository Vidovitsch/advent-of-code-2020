package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadPuzzleInput(path string) []string {
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

type Password struct {
	Policy string
	Value  string
}

func CreatePassword(passwordEntry string) *Password {
	s := strings.Split(passwordEntry, ": ")

	return &Password{Policy: s[0], Value: s[1]}
}

func ToInt(s string) int {
	result, _ := strconv.Atoi(s)

	return result
}

func (password *Password) IsValid() bool {
	p := strings.Split(password.Policy, " ")
	indices, policyChar := strings.Split(p[0], "-"), p[1]
	index1, index2 := ToInt(indices[0]), ToInt(indices[1])

	char1, char2 := string(password.Value[index1-1]), string(password.Value[index2-1])

	return char1 != char2 && (char1 == policyChar || char2 == policyChar)
}

func Solve(passwordList []string) int {
	validPasswords := 0

	for _, entry := range passwordList {
		if CreatePassword(entry).IsValid() {
			validPasswords += 1
		}
	}
	return validPasswords
}

func main() {
	puzzleInput := ReadPuzzleInput("./../input.txt")

	solution := Solve(puzzleInput)

	fmt.Printf("%v", solution)
}
