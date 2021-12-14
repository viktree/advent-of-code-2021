package main

import (
	"advent-of-code-2021/lib"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func PartOne() {
	inputs := lib.ReadInputFile("input.txt")

	positionStrings := strings.Split(inputs[0], ",")
	position := []int{}

	var maxN int
	for _, s := range positionStrings {
		n, _ := strconv.Atoi(s)
		position = append(position, n)
		if n > maxN {
			maxN = n
		}
	}

	smallestSum := 999999999999
	for j := 0; j <= maxN; j++ {
		sum := 0
		for _, n := range position {
			sum += int(math.Abs(float64(n - j)))
		}
		if sum < smallestSum {
			smallestSum = sum
		}
	}
	fmt.Println(smallestSum)
}

func PartTwo() {
	inputs := lib.ReadInputFile("input.txt")

	positionStrings := strings.Split(inputs[0], ",")
	position := []int{}

	var maxN int
	for _, s := range positionStrings {
		n, _ := strconv.Atoi(s)
		position = append(position, n)
		if n > maxN {
			maxN = n
		}
	}

	smallestSum := 999999999999
	for j := 0; j <= maxN; j++ {
		score := 0
		for _, n := range position {
			upper := int(math.Abs(float64(n - j)))
			score += ((upper) * (upper + 1)) / 2
		}
		if score < smallestSum {
			smallestSum = score
		}
	}
	fmt.Println(smallestSum)

}

func main() {
	PartOne()
	PartTwo()
}
