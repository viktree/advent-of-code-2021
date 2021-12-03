package main

import (
	"advent-of-code-2021/lib"
	"strconv"
)

func PartOne() {

	inputs := lib.ReadInputFile("input.txt")

	previous, _ := strconv.Atoi(inputs[0])
	increases := 0
	for _, input := range inputs {
		current, _ := strconv.Atoi(input)
		if current > previous {
			increases = increases + 1
		}
		previous = current
	}
	print(increases, "\n")
}

func PartTwo() {
	inputs := lib.ReadInputFile("input.txt")

	previousSum := 0
	increases := 0

	for i := 3; i < len(inputs); i++ {
		a, _ := strconv.Atoi(inputs[i-2])
		b, _ := strconv.Atoi(inputs[i-1])
		c, _ := strconv.Atoi(inputs[i])

		currentSum := a + b + c

		if currentSum > previousSum {
			increases++
		}

		previousSum = currentSum

	}
	println(increases, "\n")

}

func main() {
	PartOne()

	PartTwo()
}
