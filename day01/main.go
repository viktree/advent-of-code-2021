package main

import (
	"advent-of-code-2021/lib"
)

func PartOne() {

	inputs := lib.ReadInputFileToIntegers("input.txt")

	previous := inputs[0]
	increases := 0
	for _, current := range inputs {
		if current > previous {
			increases = increases + 1
		}
		previous = current
	}
	print(increases, "\n")
}

func PartTwo() {
	inputs := lib.ReadInputFileToIntegers("input.txt")

	previousSum := 0
	increases := 0

	for i := 3; i < len(inputs); i++ {
		currentSum := inputs[i] + inputs[i-1] + inputs[i-2]

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
