package main

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
	"math"
	"slices"
	"strconv"
	"strings"
	"utils"
)

type Solve interface {
	solve([]Cards) int
}
type Cards struct {
	no     int
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

func extractCardNo(s string) int {
	card := strings.Split(s, "Card")
	number, _ := strconv.Atoi(strings.TrimSpace(card[1]))
	return number
}

func parse(lines []string) []Cards {
	var card []Cards
	for _, line := range lines {
		game := strings.Split(line, ": ")
		if len(game) != 1 {
			id := extractCardNo(game[0])
			numbers := strings.Split(game[1], "|")
			arr1, arr2 := extractInt(numbers)
			card = append(card, Cards{id, arr1, arr2})
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

func (_ Solution1) solve(cards []Cards) int {
	total := 0
	for _, card := range cards {
		intersection := card.Intersection()
		total += double(len(intersection))
	}
	return total
}

func (_ Solution2) solve(cards []Cards) int {
	result := 0
	gameStack := stack.New()
	points := map[int]int{}
	commons := map[int]int{}

	for _, card := range cards {
		common := card.Intersection()
		points[card.no] = 0
		commons[card.no] = len(common)
		gameStack.Push(card.no)
	}

	for gameStack.Len() > 0 {
		card := gameStack.Pop().(int)
		points[card] = points[card] + 1
		common := commons[card]
		for i := 1; i <= common; i++ {
			gameStack.Push(card + i)
		}
	}

	for _, point := range points {
		result = point + result
	}
	return result
}

func parseAndSolve[T Solve](problem T, filename string) int {
	lines := utils.ReadLine(filename)
	cards := parse(lines)

	return problem.solve(cards)
}
func main() {
	fmt.Printf("Solution1: %d\nSolution2: %d", parseAndSolve(Solution1{}, "./day4/input.txt"), parseAndSolve(Solution2{}, "./day4/input.txt"))
}
