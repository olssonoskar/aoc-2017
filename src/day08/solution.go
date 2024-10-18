package day08

import (
	"aoc2017/src/collections"
	"aoc2017/src/util"
	"fmt"
	"strconv"
	"strings"
)

type pc map[string]int

type instruction struct {
	register  string
	operation string
	value     int
	condition condition
}

type condition struct {
	register string
	check    string
	value    int
}

const max = "pcMax"

func Solve() {
	input := util.GetLines("day08", "\n")
	p := make(pc, 0)
	for _, line := range input {
		inst := parse(line)
		p.next(inst)
	}
	// Take 2nd element as the first will be the 'max' of the run
	// Could be that these would be the same, then we would need to do a smarter check here
	fmt.Println("Part 1: ", collections.SortDesc(p)[1])
	fmt.Println("Part 2: ", p[max])
}

func parse(line string) *instruction {
	parts := strings.Split(strings.ReplaceAll(line, "\r", ""), " ")
	first, err1 := strconv.Atoi(parts[2])
	second, err2 := strconv.Atoi(parts[6])
	if err1 != nil || err2 != nil {
		fmt.Println("Unable to parse line: ", line)
		return nil
	}
	return &instruction{
		register:  parts[0],
		operation: parts[1],
		value:     first,
		condition: condition{
			register: parts[4],
			check:    parts[5],
			value:    second,
		},
	}
}

func (p pc) next(inst *instruction) {
	execute := operation(inst.condition, p[inst.condition.register])
	if execute {
		if inst.operation == "inc" {
			p[inst.register] += inst.value
		} else {
			p[inst.register] -= inst.value
		}
		if p[inst.register] > p[max] {
			p[max] = p[inst.register]
		}
	}
}

func operation(cond condition, regVal int) bool {
	switch cond.check {
	case "==":
		return regVal == cond.value
	case "!=":
		return regVal != cond.value
	case ">=":
		return regVal >= cond.value
	case "<=":
		return regVal <= cond.value
	case ">":
		return regVal > cond.value
	case "<":
		return regVal < cond.value
	}
	fmt.Println("Did not match a valid operation ", cond.check)
	return false
}
