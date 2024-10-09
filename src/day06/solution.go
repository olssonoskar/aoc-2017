package day06

import (
	"aoc2017/src/util"
	"fmt"
	"strconv"
	"strings"
)

func Solve() {
	input := strings.Split(util.GetInput("day06"), " ")
	banks := make([]int, len(input))
	for idx, val := range input {
		num, _ := strconv.Atoi(val)
		banks[idx] = num
	}
	total, loop := findDuplicate(banks)
	println("Part 1 = " + fmt.Sprint(total))
	println("Part 2 = " + fmt.Sprint(loop))
}

func findDuplicate(banks []int) (int, int) {
	iter := 0
	seenBanks := make(map[string]int)
	seenBanks[toString(banks)] = 0
	for {
		iter++
		banks = distribute(banks)
		result := toString(banks)
		if seenAt, ok := seenBanks[result]; ok {
			return iter, iter - seenAt
		}
		seenBanks[result] = iter
	}
}

func distribute(banks []int) []int {
	idx, blocks := blocksAtMax(banks)
	banks[idx] = 0
	for blocks > 0 {
		idx++
		banks[idx%len(banks)]++
		blocks--
	}
	return banks
}

func blocksAtMax(banks []int) (int, int) {
	indexOfMax := -1
	max := 0
	for idx, block := range banks {
		if block > max {
			max = block
			indexOfMax = idx
		}
	}
	return indexOfMax, max
}

func toString(banks []int) string {
	res := ""
	for _, block := range banks {
		current := fmt.Sprint(block)
		res += current + ","
	}
	return res[:len(res)-1]
}
