package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func part1(inp string) int {
	numbers := []string{}
	for s, r := range inp {
		if unicode.IsDigit(r) {
			numbers = append(numbers, string(inp[s]))
			break
		}
	}
	sum := numbers[0] + numbers[len(numbers)-1]
	sumint, _ := strconv.Atoi(sum)
	return sumint
}

func part2() {

}
func main() {
	input, err := os.ReadFile("./input.txt")
	if err != nil {
		panic("Fucked up")
	}
	s := string(input)
	split := strings.Fields(s)
	total := 0
	for s := range split {
		sum := part1(split[s])
		total += sum
	}
	fmt.Println(total)
}
