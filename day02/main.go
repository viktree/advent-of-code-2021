package main

import (
	"advent-of-code-2021/lib"
	"fmt"
	"strconv"
	"strings"
)

const (
	UP      = "up"
	DOWN    = "down"
	FORWARD = "forward"
)

func parseCommand(rawCmd string) (string, int) {
	cmdArr := strings.Split(rawCmd, " ")
	amount, _ := strconv.Atoi(cmdArr[1])
	return cmdArr[0], amount
}

func PartOne() {
	inputs := lib.ReadInputFile("input.txt")

	vertical, horizontal := 0, 0

	for _, input := range inputs {
		dir, x := parseCommand(input)
		if dir == DOWN {
			vertical += x
		}
		if dir == UP {
			vertical -= x
		}
		if dir == FORWARD {
			horizontal += x
		}
	}

	fmt.Println(vertical * horizontal)
}

func PartTwo() {
	inputs := lib.ReadInputFile("input.txt")

	vertical, horizontal := 0, 0
	aim := 0

	for _, input := range inputs {
		dir, x := parseCommand(input)
		if dir == DOWN {
			aim += x
		}
		if dir == UP {
			aim -= x
		}
		if dir == FORWARD {
			horizontal += x
			vertical += aim * x
		}
	}

	fmt.Println(vertical * horizontal)

}

func main() {
	PartOne()
	PartTwo()
}
