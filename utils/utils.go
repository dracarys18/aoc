package utils

import (
	"bufio"
	"os"
)

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
