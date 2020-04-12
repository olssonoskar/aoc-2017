package aoc2017

import (
	"bufio"
	"os"
)

func readFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer func() {
		file.Close()
	}()

	var content string
	scanner := bufio.NewScanner(file)
	scanner.Split

	return ""
}
