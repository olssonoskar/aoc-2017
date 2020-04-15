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

	fmt.Println("part 1: ", sum)

	sum = 0
	for _, row := range elemRows {
		quotient := findQuotient(row)
		sum += quotient
	}

	fmt.Println("part 2: ", sum)
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

func findQuotient(row []string) int {
	for i, outer := range row {
		for j, inner := range row {
			if (i != j) {
				found, quotient := isDivisible(outer, inner)
				if found {
					return quotient
				}
			}
		}
	}
	fmt.Println("No quotient found in ", row)
	return 0
}

func isDivisible(e1, e2 string) (bool, int) {
	num, err := strconv.Atoi(e1)
	if err != nil {
		panic(err)
	}
	num2, err2 := strconv.Atoi(e2)
	if err2 != nil {
		panic(err)
	}
	divident, divisor := order(num, num2)
	if (divident % divisor == 0) {
		return true, divident / divisor
	}
    return false, 0
}

func order(e1, e2 int) (high, low int) {
	if (e1 > e2) {
		return e1, e2
	}
	return e2, e1
}