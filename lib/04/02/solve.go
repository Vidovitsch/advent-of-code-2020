package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type passport struct {
	value map[string]string
}

func toInt(s string) int {
	result, _ := strconv.Atoi(s)

	return result
}

func isValidHairColor(s string) bool {
	r, _ := regexp.Compile("^#[0-9a-z]{6}$")
	return r.MatchString(s)
}

func isValidEyeColor(s string) bool {
	r, _ := regexp.Compile("^(amb|blu|brn|gry|grn|hzl|oth)$")
	return r.MatchString(s)
}

func isValidPassportId(s string) bool {
	r, _ := regexp.Compile("^\\d{9}$")
	return r.MatchString(s)
}

func isValidHeight(s string) bool {
	r, _ := regexp.Compile("^\\d{2,3}(in|cm)$")
	if !r.MatchString(s) {
		return false
	}
	height := toInt(string(s[:len(s)-2]))
	suffix := string(s[len(s)-2:])
	if suffix == "cm" {
		return height >= 150 && height <= 193
	}
	return height >= 59 && height <= 76
}

func isValidYear(s string, lowerBound int, upperBound int) bool {
	r, _ := regexp.Compile("^\\d{4}$")
	if !r.MatchString(s) {
		return false
	}
	year := toInt(s)
	return year >= lowerBound && year <= upperBound
}

func (p *passport) isValid() bool {
	for _, requiredField := range []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} {
		value, exists := p.value[requiredField]
		if !exists {
			return false
		}

		switch requiredField {
		case "byr":
			if !isValidYear(value, 1920, 2002) {
				return false
			}
		case "iyr":
			if !isValidYear(value, 2010, 2020) {
				return false
			}
		case "eyr":
			if !isValidYear(value, 2020, 2030) {
				return false
			}
		case "hgt":
			if !isValidHeight(value) {
				return false
			}
		case "hcl":
			if !isValidHairColor(value) {
				return false
			}
		case "ecl":
			if !isValidEyeColor(value) {
				return false
			}
		case "pid":
			if !isValidPassportId(value) {
				return false
			}
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
