package main

import (
	"advent-of-code-2021/lib"
	"fmt"
)

type Point struct {
	x int
	y int
}

type Move struct {
	p1 Point
	p2 Point
	dY int
	dX int
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func PartOne() {

	inputs := lib.ReadInputFile("input.txt")

	var horizontalMoves, verticalMoves []Move
	var m Move
	for _, rawMove := range inputs {
		fmt.Sscanf(rawMove, "%d,%d -> %d,%d", &m.p1.x, &m.p1.y, &m.p2.x, &m.p2.y)

		if m.p2.y < m.p1.y {
			m.dY = -1
		} else if m.p2.y > m.p1.y {
			m.dY = 1
		} else {
			m.dY = 0
		}

		if m.p2.x < m.p1.x {
			m.dX = -1
		} else if m.p2.x > m.p1.x {
			m.dX = 1
		} else {
			m.dX = 0
		}

		if m.p1.y == m.p2.y {
			horizontalMoves = append(horizontalMoves, m)
		}
		if m.p1.x == m.p2.x {
			verticalMoves = append(verticalMoves, m)
		}

	}

	visitedTwice := map[Point]bool{}

	for _, move := range horizontalMoves {
		start, stop := move.p1.x, move.p2.x
		if stop < start {
			start, stop = stop, start
		}
		for i := start; i <= stop; i++ {
			p := Point{i, move.p1.y}
			_, found := visitedTwice[p]
			visitedTwice[p] = found
		}
	}

	for _, move := range verticalMoves {
		start, stop := move.p1.y, move.p2.y
		if stop < start {
			start, stop = stop, start
		}
		for i := start; i <= stop; i++ {
			p := Point{move.p1.x, i}
			_, found := visitedTwice[p]
			visitedTwice[p] = found
		}
	}

	count := 0
	for _, v := range visitedTwice {
		if v {
			count++
		}
	}

	fmt.Println(count)
}

func PartTwo() {
	inputs := lib.ReadInputFile("input.txt")

	var moves []Move
	var m Move
	for _, rawMove := range inputs {
		fmt.Sscanf(rawMove, "%d,%d -> %d,%d", &m.p1.x, &m.p1.y, &m.p2.x, &m.p2.y)

		if m.p2.y < m.p1.y {
			m.dY = -1
		} else if m.p2.y > m.p1.y {
			m.dY = 1
		} else {
			m.dY = 0
		}

		if m.p2.x < m.p1.x {
			m.dX = -1
		} else if m.p2.x > m.p1.x {
			m.dX = 1
		} else {
			m.dX = 0
		}
		moves = append(moves, m)

	}

	visitedTwice := map[Point]bool{}

	for _, move := range moves {
		x, y := move.p1.x, move.p1.y

		for {
			_, found := visitedTwice[Point{x, y}]
			visitedTwice[Point{x, y}] = found

			if x == move.p2.x && y == move.p2.y {
				break
			}
			x, y = x+move.dX, y+move.dY
		}
	}

	count := 0
	for _, v := range visitedTwice {
		if v {
			count++
		}
	}

	fmt.Println(count)

}

func main() {
	PartOne()

	PartTwo()
}
