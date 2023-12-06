package utils

import (
	"bufio"
	"errors"
	"os"
	"runtime"
	"strconv"
	"sync"
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

func ProductArray(array []int) int {
	product := 1
	for _, ele := range array {
		product *= ele
	}
	return product
}

func AtoiArr(array []string) []int {
	var result []int
	for _, s := range array {
		num, err := strconv.Atoi(s)
		if err == nil {
			result = append(result, num)
		}
	}
	return result
}

func MapParallel[T any, B any](array []T, inner func(T) B) []B {
	//Change the number of cores needed to max cores available
	runtime.GOMAXPROCS(runtime.NumCPU())

	var result []B
	waitgroup := sync.WaitGroup{}

	channel := make(chan B, len(array))

	waitgroup.Add(len(array))

	do := func(val T) {
		res := inner(val)
		channel <- res
		waitgroup.Done()
	}
	for _, proc := range array {
		go do(proc)
	}

	waitgroup.Wait()
	close(channel)

	for ele := range channel {
		result = append(result, ele)
	}

	return result
}

func ReadString(filename string) string {
	content, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return string(content)
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
