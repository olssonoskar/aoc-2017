package day03

type Direction int

const (
	R Direction = iota
	U
	L
	D
)

type Battery [][]int
type Position struct {
	x, y int
	dir  Direction
}

func Part2() {
	pos := &Position{10, 10, D}
	battery := initBattery()
	battery[pos.x][pos.y] = 1
	for {
		turnIfPossible(battery, pos)
		sum := battery.sumNeighbors(pos)
		if sum > 312051 {
			println(sum)
			return
		}
		battery[pos.x][pos.y] = sum
	}
}

func (b Battery) sumNeighbors(pos *Position) int {
	sum := 0
	for x := pos.x - 1; x <= pos.x+1; x++ {
		for y := pos.y - 1; y <= pos.y+1; y++ {
			if x == pos.x && y == pos.y {
				continue
			}
			sum += b[x][y]
		}
	}
	return sum
}

// Attemps turn, if not free continue to next
func turnIfPossible(battery [][]int, pos *Position) {
	switch pos.dir {
	case D:
		res := battery[pos.x+1][pos.y]
		if res == 0 {
			pos.x += 1
			pos.dir = R
		} else {
			pos.y -= 1
		}
	case R:
		res := battery[pos.x][pos.y+1]
		if res == 0 {
			pos.y += 1
			pos.dir = U
		} else {
			pos.x += 1
		}
	case U:
		res := battery[pos.x-1][pos.y]
		if res == 0 {
			pos.x -= 1
			pos.dir = L
		} else {
			pos.y += 1
		}
	case L:
		res := battery[pos.x][pos.y-1]
		if res == 0 {
			pos.y -= 1
			pos.dir = D
		} else {
			pos.x -= 1
		}
	}
}

func initBattery() Battery {
	battery := make([][]int, 20)
	for i := range battery {
		battery[i] = make([]int, 20)
	}
	return battery
}
