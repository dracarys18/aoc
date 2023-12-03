package utils

import (
	"bufio"
	"errors"
	"os"
)

type Adder interface {
	int | float64 | string
}

func Get[T any](array []T, index int) (T, error) {
	var val T
	err := errors.New("Index out of bounds")
	if index < len(array) || index >= 0 {
		val = array[index]
		err = nil
	}
	return val, err
}

func SumArray[T Adder](array []T) T {
	var sum T
	for _, ele := range array {
		sum += ele
	}
	return sum
}

func ReadLine(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scan := bufio.NewScanner(file)
	lines := []string{}
	for scan.Scan() {
		line := scan.Text()
		lines = append(lines, line)
	}
	file.Close()
	return lines
}
