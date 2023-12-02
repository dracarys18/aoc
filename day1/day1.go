package main

import (
	"fmt"
	"strconv"
	"unicode"
	"utils"
)

func part1(inp string) int {
	numbers := []string{}
	for _, r := range inp {
		if unicode.IsDigit(r) {
			numbers = append(numbers, string(r))
		}
	}
	sum := "0"
	if len(numbers) != 0 {
		sum = numbers[0] + numbers[len(numbers)-1]
	}
	sumint, _ := strconv.Atoi(sum)
	return sumint
}

func part2() {

}
func main() {
	lines := utils.ReadLine("./day1/input.txt")
	total := 0
	for _, line := range lines {
		sum := part1(line)
		total += sum
	}
	fmt.Println(total)
}
