package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {
	dat, err := ioutil.ReadFile("day1.txt")
	if err != nil {
		panic(err)
	}
	input := string(dat)

	fmt.Println("part 1 =", Part1(input))
	fmt.Println("part 2 =", Part2(input))
}

func Part1(input string) int {
	return review(input, 1)
}

func Part2(input string) int {
	return review(input, len(input)/2)
}

func review(input string, offset int) (sum int) {
	sum = 0
	inputLength := len(input)
	for index := 0; index < inputLength; index++ {
		current := input[index]
		refrence := input[(index+offset)%inputLength]
		if current == refrence {
			num, _ := strconv.Atoi(string(current))
			sum += num
		}
	}
	return
}
