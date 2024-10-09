package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// GetInput of specified day
func GetInput(folder string) string {
	filepath, err := filepath.Abs("../src/" + folder + "/input.txt")
	if err != nil {
		fmt.Print(err)
	}
	input, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Print(err)
	}
	return string(input)
}

func GetLines(folder, delimiter string) []string {
	return strings.Split(GetInput(folder), delimiter)
}
