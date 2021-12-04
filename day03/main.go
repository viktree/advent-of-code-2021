package main

import (
	"advent-of-code-2021/lib"
	"fmt"
	"sort"
	"strconv"
)

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func PartOne() {

	inputs := lib.ReadInputFile("input.txt")

	sizeOfEntries := len(inputs[0])

	count := make([]int, sizeOfEntries)
	gammaString := make([]byte, sizeOfEntries)
	epsilonString := make([]byte, sizeOfEntries)

	for pos := range count {
		count[pos] = 0
	}

	for _, input := range inputs {
		for pos, i := range input {
			if i == '0' {
				count[pos]++
			}
		}
	}

	for pos := range count {
		if count[pos] < (len(inputs) / 2) {
			gammaString[pos] = '0'
			epsilonString[pos] = '1'
		} else {
			gammaString[pos] = '1'
			epsilonString[pos] = '0'
		}
	}

	gamma, _ := strconv.ParseUint(string(gammaString), 2, sizeOfEntries)
	epsilon, _ := strconv.ParseUint(string(epsilonString), 2, sizeOfEntries)

	fmt.Println(gamma * epsilon)
}

func PartTwo() {

	inputs := lib.ReadInputFile("input.txt")
	sort.Strings(inputs)

	sizeOfEntries := len(inputs[0])

	prevScrubberSearchSpace := inputs
	nextScrubberSearchSpace := []string{}
	for i := 0; i < sizeOfEntries; i++ {
		numberOfOnes := 0
		numberOfZeros := 0
		for _, input := range prevScrubberSearchSpace {
			if input[i] == '1' {
				numberOfOnes++
			} else {
				numberOfZeros++
			}
		}
		scrubberSearchVal := byte('1')
		if numberOfOnes < numberOfZeros {
			scrubberSearchVal = byte('0')
		}
		for _, input := range prevScrubberSearchSpace {
			if input[i] == scrubberSearchVal {
				nextScrubberSearchSpace = append(nextScrubberSearchSpace, input)
			}
		}
		if len(nextScrubberSearchSpace) == 1 {
			break
		}
		prevScrubberSearchSpace = nextScrubberSearchSpace
		nextScrubberSearchSpace = []string{}
	}

	prevGeneratorSearchSpace := inputs
	nextGeneratorSearchSpace := []string{}
	for i := 0; i < sizeOfEntries; i++ {
		numberOfOnes := 0
		numberOfZeros := 0
		for _, input := range prevGeneratorSearchSpace {
			if input[i] == '1' {
				numberOfOnes++
			} else {
				numberOfZeros++
			}
		}
		scrubberSearchVal := byte('0')
		if numberOfOnes < numberOfZeros {
			scrubberSearchVal = byte('1')
		}
		for _, input := range prevGeneratorSearchSpace {
			if input[i] == scrubberSearchVal {
				nextGeneratorSearchSpace = append(nextGeneratorSearchSpace, input)
			}
		}
		if len(nextGeneratorSearchSpace) == 1 {
			break
		}
		prevGeneratorSearchSpace = nextGeneratorSearchSpace
		nextGeneratorSearchSpace = []string{}
	}

	generatorRating, _ := strconv.ParseUint(nextGeneratorSearchSpace[0], 2, sizeOfEntries)
	scrubberRating, _ := strconv.ParseUint(nextScrubberSearchSpace[0], 2, sizeOfEntries)

	fmt.Println(generatorRating * scrubberRating)

}

func main() {
	PartOne()

	PartTwo()
}
