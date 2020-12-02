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

func solve(expenseReport []int, target int) (int, error) {
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

func main() {
	expenseReport := readExpenseReport("./../input.txt")

	solution, err := solve(expenseReport, 2020)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Solution:", solution) // 1010884
}
