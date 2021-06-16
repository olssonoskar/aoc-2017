package main

func main() {
	//part1()
	Part2()
}

func part1() {
	seeked := 312051
	upper, lower, layer := getCircleRange(seeked)
	println("Initial", upper, lower)
	toMiddle := stepsFromMiddle(seeked, upper, lower)
	println(toMiddle)
	println((layer - 1) + toMiddle)
}

func getCircleRange(value int) (lower, upper, index int) {
	upper = 1
	lower = 1
	for index := 0; true; index++ {
		if value <= upper && value >= lower {
			return upper, lower, index
		}
		lower = upper + 1
		upper = circleRange(sideLength(index), upper)
	}
	return upper, lower, index
}

func sideLength(index int) int {
	if index == 1 {
		return 1
	}
	return (index * 2) - 1
}

func circleRange(sideLength, previousRange int) int {
	return (sideLength * 4) + 4 + previousRange
}

func stepsFromMiddle(seeked, upper, lower int) int {
	from := lower - 1
	sideLemgth := ((upper - lower + 1) / 4)
	to := from + sideLemgth
	for i := 1; i < 5; i++ {
		println("Searching between", from, " ", to, "sidelength", sideLemgth)
		if seeked >= from && seeked <= to {
			return fromMiddle(seeked, from+(sideLemgth/2))
		}
		from += sideLemgth
		to += sideLemgth
	}
	return -1
}

func fromMiddle(seeked, middle int) int {
	println("expected middle", middle)
	if seeked > middle {
		return seeked - middle
	}
	return middle - seeked
}
