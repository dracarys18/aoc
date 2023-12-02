package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var possibility = map[string]int{
	"red":   12,
	"blue":  14,
	"green": 13,
}

func solution1(game string, index int) bool {
	tries := strings.Split(game, ";")
	for _, try := range tries {
		subsets := strings.Split(try, ",")
		actual := map[string]int{
			"red":   0,
			"blue":  0,
			"green": 0,
		}
		for _, subset := range subsets {
			play := strings.Split(subset, " ")
			count, err := strconv.Atoi(play[1])
			if err == nil {
				actual[play[2]] += count
			}
		}
		for k, v := range actual {
			if v > possibility[k] {
				return false
			}
		}
	}
	return true
}

func parseAndSolve(filename string) int {
	input, err := os.ReadFile(filename)
	if err != nil {
		panic("Unable to read the file")
	}
	lines := strings.Split(string(input), "\n")
	sum := 0
	for _, s := range lines {
		game := strings.Split(s, ":")
		if len(game[0]) != 0 {
			gameIndex, err := strconv.Atoi(game[0][5:])
			if err != nil {
				panic("Invalid input")
			}
			if solution1(game[1], gameIndex) {
				sum += gameIndex
			}
		}
	}
	return sum
}

func main() {
	fmt.Println(parseAndSolve("input.txt"))
}
