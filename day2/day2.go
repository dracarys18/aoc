package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"utils"
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

func solution2(game string) int {
	mintries := map[string]int{
		"red":   math.MinInt64,
		"blue":  math.MinInt64,
		"green": math.MinInt64,
	}

	tries := strings.Split(game, ";")
	for _, try := range tries {
		subsets := strings.Split(try, ",")
		for _, subset := range subsets {
			play := strings.Split(subset, " ")
			count, err := strconv.Atoi(play[1])
			if err == nil {
				mintries[play[2]] = max(mintries[play[2]], count)
			}
		}
	}
	power := 1
	for _, v := range mintries {
		power *= v
	}
	return power
}

func parseAndSolve(filename string, which string) int {
	lines := utils.ReadLine(filename)
	sum := 0
	for _, s := range lines {
		game := strings.Split(s, ":")
		if len(game[0]) != 0 {
			gameIndex, err := strconv.Atoi(game[0][5:])
			if err != nil {
				panic("Invalid input")
			}
			switch which {
			case "sol1":
				if solution1(game[1], gameIndex) {
					sum += gameIndex
				}
			case "sol2":
				sum += solution2(game[1])
			default:
				panic("Wrong solution number")

			}
		}
	}
	return sum
}

func main() {
	fmt.Printf("sol1 %d\n", parseAndSolve("./day2/input.txt", "sol1"))
	fmt.Printf("sol2 %d\n", parseAndSolve("./day2/input.txt", "sol2"))
}
