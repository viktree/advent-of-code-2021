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

const (
	FORWARD  = 1
	BACKWARD = -1
	UP       = -1
	DOWN     = 1
	STEADY   = 0
)

func parseMove(rawMove string) (m Move) {
	fmt.Sscanf(rawMove, "%d,%d -> %d,%d", &m.p1.x, &m.p1.y, &m.p2.x, &m.p2.y)

	if m.p2.y < m.p1.y {
		m.dY = UP
	} else if m.p2.y > m.p1.y {
		m.dY = DOWN
	} else {
		m.dY = STEADY
	}

	if m.p2.x < m.p1.x {
		m.dX = BACKWARD
	} else if m.p2.x > m.p1.x {
		m.dX = FORWARD
	} else {
		m.dX = STEADY
	}

	return m
}

func performMoves(moves []Move) (visitedTwice map[Point]bool) {
	visitedTwice = make(map[Point]bool, 0)
	for _, move := range moves {
		p := move.p1
		for {
			_, found := visitedTwice[p]
			visitedTwice[p] = found

			if p == move.p2 {
				break
			}
			p = Point{p.x + move.dX, p.y + move.dY}
		}
	}
	return visitedTwice
}

func countEntries(set map[Point]bool) (count int) {
	count = 0
	for _, v := range set {
		if v {
			count++
		}
	}
	return count
}

func PartOne() {
	inputs := lib.ReadInputFile("input.txt")

	var moves []Move
	for _, rawMove := range inputs {
		m := parseMove(rawMove)
		if m.p1.y == m.p2.y || m.p1.x == m.p2.x {
			moves = append(moves, m)
		}
	}

	visitedTwice := performMoves(moves)
	count := countEntries(visitedTwice)

	fmt.Println(count)
}

func PartTwo() {
	inputs := lib.ReadInputFile("input.txt")

	var moves []Move
	for _, rawMove := range inputs {
		m := parseMove(rawMove)
		moves = append(moves, m)
	}

	visitedTwice := performMoves(moves)
	count := countEntries(visitedTwice)

	fmt.Println(count)
}

func main() {
	PartOne()
	PartTwo()
}
