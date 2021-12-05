package main

import (
	"advent-of-code-2021/lib"
	"fmt"
)

func PartOne() {
	inputs := lib.ReadInputFileToIntegers("input.txt")

	previous, increases := inputs[0], 0

	for _, current := range inputs {
		if current > previous {
			increases++
		}
		previous = current
	}
	fmt.Println(increases)
}

func PartTwo() {
	inputs := lib.ReadInputFileToIntegers("input.txt")

	previousSum, increases := 0, 0

	for i := 3; i < len(inputs); i++ {
		currentSum := inputs[i] + inputs[i-1] + inputs[i-2]

		if currentSum > previousSum {
			increases++
		}

		previousSum = currentSum

	}
	fmt.Println(increases)

}

func main() {
	PartOne()
	PartTwo()
}
