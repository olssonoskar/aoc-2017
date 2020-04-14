package main

import (
	"io/ioutil"
	"strconv"
	"fmt"
	"strings"
	"math"
)

func main() {
	dat, err := ioutil.ReadFile("day2.txt")
	if err != nil {
		panic(err)
	}

	rows := strings.Split(string(dat), "\r\n")
	var elemRows [][]string = make([][]string, len(rows))

	sum := 0
	for i, row := range rows {
		elemRows[i] = strings.Split(row, "\t")
		min, max := findMinMax(elemRows[i])
		sum += (max - min)
	}

	fmt.Println(sum)
}

func findMinMax(row []string) (min, max int) {
	min = math.MaxInt32
	max = 0
	for _, elem := range row {
		num, err := strconv.Atoi(elem)
		if err != nil {
			panic(err)
		}
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return
}

func findFirstDivisor(row []string) (divident, divisor int) {
	//TODO
	return
}