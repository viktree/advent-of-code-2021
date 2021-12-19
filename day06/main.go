package main

import (
	"advent-of-code-2021/lib"
	"fmt"
	"strconv"
	"strings"
)

func modLikePython(d, m int) int {
	var res int = d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
	}
	return res
}

func parseFish(input string) (fish []int) {
	fishStrings := strings.Split(input, ",")
	fish = make([]int, 0)
	for _, f := range fishStrings {
		fInt, _ := strconv.Atoi(f)
		fish = append(fish, fInt)
	}
	return fish
}

func updateFish(curr int) int {
	if curr == 8 {
		return 7
	}
	if curr == 7 {
		return 6
	}
	return modLikePython(curr-1, 7)
}

func sum(arr [9]int) (total int) {
	for _, item := range arr {
		total += item
	}
	return total
}

func PartOne() {
	inputs := lib.ReadInputFile("input.txt")
	fishPreviousDay := parseFish(inputs[0])
	var totalFishToday int
	for day := 1; day <= 80; day++ {
		fishToday := []int{}
		newFishCount := 0
		for _, f := range fishPreviousDay {
			updatedFish := updateFish(f)
			fishToday = append(fishToday, updatedFish)
			if f == 0 {
				newFishCount++
			}
		}
		for i := 0; i < newFishCount; i++ {
			fishToday = append(fishToday, 8)
		}
		totalFishToday = len(fishToday)
		fishPreviousDay = fishToday
	}

	fmt.Println(totalFishToday)
}

func PartTwo() {
	inputs := lib.ReadInputFile("simple.txt")

	var fishCycles [9]int
	fishPreviousDay := parseFish(inputs[0])

	for _, f := range fishPreviousDay {
		fishCycles[f]++
	}

	var totalFishToday int
	var newFishCycles [9]int
	for day := 1; day <= 80; day++ {
		newFishCycles[0] = fishCycles[0]
		newFishCycles[1] = fishCycles[1]
		newFishCycles[2] = fishCycles[2]
		newFishCycles[3] = fishCycles[3]
		newFishCycles[4] = fishCycles[4]
		newFishCycles[5] = fishCycles[5]
		newFishCycles[6] = fishCycles[6]
		newFishCycles[7] = fishCycles[7]
		newFishCycles[8] = fishCycles[8]

		totalFishToday = sum(newFishCycles)
		fishCycles = newFishCycles
		fmt.Println(day, totalFishToday, fishCycles)
	}

	fmt.Println(totalFishToday)

}

func main() {
	PartOne()
	PartTwo()
}
