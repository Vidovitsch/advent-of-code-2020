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
	expenseReport := readExpenseReport("./../input.txt")

	solution, err := solve(expenseReport, 2020)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Solution:", solution) // 1010884
}
