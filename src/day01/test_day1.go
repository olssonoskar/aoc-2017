package day01

import (
	"testing"
)

func TestPart1(t *testing.T) {
	var tests = []struct {
		input   string
		expects int
	}{
		{"1122", 3},
		{"111", 3},
		{"91212129", 9},
	}
	for _, test := range tests {
		result := Part1(test.input)
		if result != test.expects {
			t.Errorf("Exptected %d, got %d", test.expects, result)
		}
	}
}

func TestPart2(t *testing.T) {
	var tests = []struct {
		input   string
		expects int
	}{
		{"1212", 6},
		{"123425", 4},
		{"12131415", 4},
	}
	for _, test := range tests {
		result := Part2(test.input)
		if result != test.expects {
			t.Errorf("Exptected %d, got %d", test.expects, result)
		}
	}
}
