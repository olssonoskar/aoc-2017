package day09

import (
	"aoc2017/src/util"
	"fmt"
)

func Solve() {
	groups := util.GetInput("day09")
	part1, part2 := Iterate(groups)
	fmt.Println("Part 1: ", part1)
	fmt.Println("Part 2: ", part2)
}

func Iterate(line string) (int, int) {
	sum := 0
	level := 0
	trashCount := 0
	trash := false
	skip := false

	for _, c := range line {
		if skip {
			skip = false
			continue
		}
		if c == '!' {
			skip = true
			continue
		}
		if c == '<' && !trash {
			trash = true
			continue
		}
		if c == '>' {
			trash = false
			continue
		}
		// Count the trash rune and continue
		if trash {
			trashCount++
			continue
		}
		// Record if end of group, Inc/Dec level
		if c == '{' {
			level += 1
		}
		if c == '}' {
			sum += level
			level -= 1
		}
	}
	return sum, trashCount
}
