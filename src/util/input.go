package util

import (
	"io/ioutil"
	"path/filepath"
)

// GetInput of specified day
func GetInput(folder string) string {
	filepath, _ := filepath.Abs("../src/" + folder + "/input.txt")
	input, _ := ioutil.ReadFile(filepath)
	return string(input)
}
