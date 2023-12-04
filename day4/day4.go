package main

import (
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"utils"
)

type Solve interface {
	solve([]string) int
}
type Cards struct {
	array1 []int
	array2 []int
}

type Solution1 struct{}
type Solution2 struct{}

func (game Cards) Intersection() []int {
	var interesection []int
	for _, ele := range game.array1 {
		if slices.Contains(game.array2, ele) {
			interesection = append(interesection, ele)
		}
	}
	return interesection
}

func MaptoInt(arr string) []int {
	s := strings.Split(arr, " ")
	var result []int
	for _, s := range s {
		n, err := strconv.Atoi(s)
		if err == nil {
			result = append(result, n)
		}
	}
	return result
}

func extractInt(numbers []string) ([]int, []int) {
	arr1 := MaptoInt(numbers[0])
	arr2 := MaptoInt(numbers[1])
	return arr1, arr2
}
func parse(lines []string) []Cards {
	var card []Cards
	for _, line := range lines {
		game := strings.Split(line, ": ")
		if len(game) != 1 {
			numbers := strings.Split(game[1], "|")
			arr1, arr2 := extractInt(numbers)
			card = append(card, Cards{array1: arr1, array2: arr2})
		}
	}
	return card
}

func double(size int) int {
	result := size
	if size > 1 {
		result = int(math.Pow(float64(2), float64(size-1)))
	}
	return result
}

func (_ Solution1) solve(lines []string) int {
	cards := parse(lines)
	total := 0
	for _, card := range cards {
		intersection := card.Intersection()
		total += double(len(intersection))
	}
	return total
}

func parseAndSolve[T Solve](problem T, filename string) int {
	lines := utils.ReadLine(filename)
	return problem.solve(lines)
}
func main() {
	fmt.Printf("Solution1: %d\n", parseAndSolve(Solution1{}, "./day4/input.txt"))
}
