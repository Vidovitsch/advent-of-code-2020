package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readExpenseReport(path string) []int {
	expenseReport := []int{}

	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		expense, _ := strconv.Atoi(scanner.Text())
		expenseReport = append(expenseReport, expense)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return expenseReport
}

func solvePart1(expenseReport []int, target int) (int, error) {
	for i, expense1 := range expenseReport {
		rest := target - expense1

		for _, expense2 := range expenseReport[i+1:] {
			if expense2 == rest {
				return expense1 * expense2, nil
			}
		}
	}
	return 0, errors.New("Part 1 is unsolvable.")
}

func solvePart2(expenseReport []int, target int) (int, error) {
	for i, expense1 := range expenseReport {

		for j, expense2 := range expenseReport[i+1:] {
			rest := target - expense1 - expense2

			for _, expense3 := range expenseReport[j+1:] {
				if expense3 == rest {
					return expense1 * expense2 * expense3, nil
				}
			}
		}
	}
	return 0, errors.New("Part 2 is unsolvable.")
}

func main() {
	expenseReport := readExpenseReport("./day1_input.txt")

	solutionPart1, err := solvePart1(expenseReport, 2020)

	if err != nil {
		log.Fatal(err)
	}

	solutionPart2, err := solvePart2(expenseReport, 2020)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Solution part 1:", solutionPart1) // 1010884
	fmt.Println("Solution part 2:", solutionPart2) // 253928438
}
