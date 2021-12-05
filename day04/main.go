package main

import (
	"advent-of-code-2021/lib"
	"fmt"
	"strconv"
	"strings"
)

type LineType bool

const (

	// Enum of two values
	Row = true
	Col = false

	// Parsing size
	PARSE_BOARD_SIZE = 6
	BOARD_SIZE       = PARSE_BOARD_SIZE - 1
)

type BingoCardLine struct {
	cardId   int
	lineType LineType
	lineIdx  int
}

func updateBoardScores(scores map[BingoCardLine]int, cardId int, amount int) map[BingoCardLine]int {
	for i := 0; i < 5; i++ {
		scores[BingoCardLine{cardId, Row, i}] += amount
		scores[BingoCardLine{cardId, Col, i}] += amount
	}
	return scores
}

func PartOne() {
	inputs := lib.ReadInputFile("input.txt")

	// Start with empty rows and cols. Need 5 markings to continue
	numberOfBoards := len(inputs) / PARSE_BOARD_SIZE

	squaresLeftInLine := map[BingoCardLine]int{}
	cardScoreForLine := map[BingoCardLine]int{}

	for cardIdx := 0; cardIdx < numberOfBoards; cardIdx++ {
		for i := 0; i < BOARD_SIZE; i++ {
			squaresLeftInLine[BingoCardLine{cardIdx, Row, i}] = BOARD_SIZE
			squaresLeftInLine[BingoCardLine{cardIdx, Col, i}] = BOARD_SIZE

			cardScoreForLine[BingoCardLine{cardIdx, Row, i}] = 0
			cardScoreForLine[BingoCardLine{cardIdx, Col, i}] = 0
		}
	}

	// Set the locations of each number that gets called
	linesOfChosenNumber := map[int]([]BingoCardLine){}
	i := 2
	cardIdx := 0
	for i < len(inputs) {
		for r := 0; r < BOARD_SIZE; r++ {
			row := strings.Fields(inputs[i])
			for c := 0; c < BOARD_SIZE; c++ {
				n := row[c]
				amount, _ := strconv.Atoi(n)

				linesOfChosenNumber[amount] = append(
					linesOfChosenNumber[amount],
					BingoCardLine{cardIdx, Row, r},
				)
				linesOfChosenNumber[amount] = append(
					linesOfChosenNumber[amount],
					BingoCardLine{cardIdx, Col, c},
				)
				cardScoreForLine = updateBoardScores(cardScoreForLine, cardIdx, amount)
			}
			i++
		}
		i++
		cardIdx++
	}

	// Let's play!!
	bingoScore := 0
	numbers := strings.Split(inputs[0], ",")

GAME_PLAY:
	for _, n := range numbers {
		amount, _ := strconv.Atoi(n)
		lines := linesOfChosenNumber[amount]

		for _, line := range lines {
			if line.lineType == Row {
				cardScoreForLine = updateBoardScores(cardScoreForLine, line.cardId, -amount)
			}
		}

		for _, line := range lines {
			squaresLeftInLine[line] -= 1
			if squaresLeftInLine[line] == 0 {
				bingoScore = amount * cardScoreForLine[line]
				break GAME_PLAY
			}
		}
	}
	fmt.Println(bingoScore)
}

func PartTwo() {
	inputs := lib.ReadInputFile("input.txt")

	// Start with empty rows and cols. Need 5 markings to continue
	numberOfBoards := len(inputs) / PARSE_BOARD_SIZE

	squaresLeftInLine := map[BingoCardLine]int{}
	cardScoreForLine := map[BingoCardLine]int{}

	for cardIdx := 0; cardIdx < numberOfBoards; cardIdx++ {
		for i := 0; i < 5; i++ {
			squaresLeftInLine[BingoCardLine{cardIdx, Row, i}] = BOARD_SIZE
			squaresLeftInLine[BingoCardLine{cardIdx, Col, i}] = BOARD_SIZE

			cardScoreForLine[BingoCardLine{cardIdx, Row, i}] = 0
			cardScoreForLine[BingoCardLine{cardIdx, Col, i}] = 0
		}
	}

	// Set the locations of each number that gets called
	linesOfChosenNumber := map[int]([]BingoCardLine){}
	i := 2
	cardIdx := 0
	for i < len(inputs) {
		for r := 0; r < BOARD_SIZE; r++ {
			row := strings.Fields(inputs[i])
			for c := 0; c < BOARD_SIZE; c++ {
				n := row[c]
				amount, _ := strconv.Atoi(n)

				linesOfChosenNumber[amount] = append(
					linesOfChosenNumber[amount],
					BingoCardLine{cardIdx, Row, r},
				)
				linesOfChosenNumber[amount] = append(
					linesOfChosenNumber[amount],
					BingoCardLine{cardIdx, Col, c},
				)

				cardScoreForLine = updateBoardScores(cardScoreForLine, cardIdx, amount)
			}
			i++
		}
		i++
		cardIdx++
	}

	// Let's play!!
	bingoScore := 0
	numbers := strings.Split(inputs[0], ",")

	cardHasBingo := map[int]bool{}

	for i := 0; i < numberOfBoards; i++ {
		cardHasBingo[i] = false
	}

	for _, n := range numbers {
		amount, _ := strconv.Atoi(n)
		lines := linesOfChosenNumber[amount]

		for _, line := range lines {
			if line.lineType == Row {
				cardScoreForLine = updateBoardScores(cardScoreForLine, line.cardId, -amount)
			}
		}

		for _, line := range lines {
			squaresLeftInLine[line] -= 1
			if squaresLeftInLine[line] == 0 && !cardHasBingo[line.cardId] {
				bingoScore = amount * cardScoreForLine[line]
				cardHasBingo[line.cardId] = true
			}
		}
	}
	fmt.Println(bingoScore)
}

func main() {
	PartOne()

	PartTwo()
}
