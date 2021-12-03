package main

import (
	"advent-of-code-2021/lib"
	"strconv"
	"strings"
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
		if dir == "down" {
			vertical += x
		}
		if dir == "up" {
			vertical -= x
		}
		if dir == "forward" {
			horizontal += x
		}
	}

	print(vertical*horizontal, "\n")
}

func PartTwo() {
	inputs := lib.ReadInputFile("input.txt")

	vertical, horizontal := 0, 0
	aim := 0

	for _, input := range inputs {
		dir, x := parseCommand(input)
		if dir == "down" {
			aim += x
		}
		if dir == "up" {
			aim -= x
		}
		if dir == "forward" {
			horizontal += x
			vertical += aim * x
		}
	}

	print(vertical*horizontal, "\n")

}

func main() {
	PartOne()
	PartTwo()
}
