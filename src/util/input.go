package util

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// GetInput of specified day
func GetInput(folder string) string {
	filepath, err := filepath.Abs("../src/" + folder + "/input.txt")
	if err != nil {
		fmt.Print(err)
	}
	input, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Print(err)
	}
	return string(input)
}

func GetLines(folder, delimiter string) []string {
	return strings.Split(GetInput(folder), delimiter)
}
