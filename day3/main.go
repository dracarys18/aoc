package main

import (
	"fmt"
	"strconv"
	"unicode"
	"utils"
)

type Solve interface {
	solve([]string) int
}

type Solution1 struct{}

type Solution2 struct{}

type Digit struct {
	digit int
	row   int
	start int
	end   int
}

func appendDigitArray(digits *[]Digit, number *string, row int, start int, end int) {
	if *number != "" {
		temp, _ := strconv.Atoi(*number)
		digit := Digit{digit: temp, row: row, start: start, end: end}

		*digits = append(*digits, digit)
		*number = ""
	}

}
func (_ Solution1) solve(lines []string) int {
	var digits []Digit
	for row, line := range lines {
		number := ""
		for col, char := range line {
			if char >= '0' && char <= '9' {
				number = number + string(char)
			} else {
				start := col - len(number)
				end := col - 1
				appendDigitArray(&digits, &number, row, start, end)
			}
		}
		// handle case when number is at the end
		end := len(line) - 1
		start := end - (len(number) - 1)
		appendDigitArray(&digits, &number, row, start, end)

	}
	total := 0
	for _, digit := range digits {
		if isAdjacent(lines, digit.row, digit.start, digit.end) {
			total += digit.digit
		}
	}

	return total
}

func isSymbol(char byte) bool {
	return char != '.' && !unicode.IsDigit(rune(char))
}

func isAdjacent(lines []string, row, start int, end int) bool {

	directions := [][2]int{
		{row - 1, start - 1},
		{row + 1, start - 1},
		{row, start - 1},
		{row, end + 1},
		{row + 1, end + 1},
		{row + 1, end - 1},
		{row - 1, end + 1},
		{row - 1, end - 1},
		{row + 1, start},
		{row + 1, end},
		{row - 1, start},
		{row - 1, end},
	}

	for _, dir := range directions {
		row, col := dir[0], dir[1]
		if row >= 0 && row < len(lines) && col >= 0 && col < len(lines[0]) {
			if isSymbol(lines[row][col]) {
				return true
			}
		}
	}

	return false
}

func (_ Solution2) solve(lines []string) int {
	return 0
}

func parseAndSolve[T Solve](filename string, sol T) int {
	lines := utils.ReadLine(filename)
	return sol.solve(lines)
}

func main() {
	fmt.Printf("solution1: %d\nsolution2: %d", parseAndSolve("./day3/input.txt", Solution1{}), parseAndSolve("./day3/input.txt", Solution2{}))
}
