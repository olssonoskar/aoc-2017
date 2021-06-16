package main

import (
	"strconv"
	"util"
)

func main() {
	input := util.GetLines("day05", "\r\n")
	println(Part1(input))
	println(Part2(input))
}

func Part1(input []string) int {
	cpu := tableFrom(input)
	return cpu.performJumps(false)
}

func Part2(input []string) int {
	cpu := tableFrom(input)
	return cpu.performJumps(true)
}

func (cpu *CPU) performJumps(offsetEnabled bool) int {
	totalJumps := 0
	for {
		jump, ok := cpu.instructions[cpu.pc]
		if !ok {
			return totalJumps
		}
		totalJumps++
		if jump >= 3 && offsetEnabled {
			cpu.instructions[cpu.pc]--
		} else {
			cpu.instructions[cpu.pc]++
		}
		cpu.pc += jump
	}
}

type CPU struct {
	pc           int
	instructions map[int]int
}

func tableFrom(input []string) *CPU {
	instrTable := make(map[int]int)
	for index, value := range input {
		number, _ := strconv.Atoi(value)
		instrTable[index] = number
	}
	return &CPU{
		pc:           0,
		instructions: instrTable,
	}
}
