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
	inputLength := len(input)

	sum := review(input, 1)
	fmt.Println("part 1 =", sum)

	sum = review(input, inputLength/2)
	fmt.Println("part 2 =", sum)
}

func parseInt(number string) int {
	val, err := strconv.Atoi(number)
	if err != nil {
		panic(err)
	}
	return val
}

func review(input string, offset int) (sum int) {
	sum = 0
	inputLength := len(input)
	for index := 0; index < inputLength; index++ {
		current := input[index]
		refrence := input[(index+offset)%inputLength]
		if current == refrence {
			sum += parseInt(string(current))
		}
	}
	return
}
