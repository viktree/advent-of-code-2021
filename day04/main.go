package main

import (
	"advent-of-code-2021/lib"
	"strconv"
	"strings"
)

type Row struct {
	board, row int
}

type Col struct {
	board, col int
}

func PartOne() {

	inputs := lib.ReadInputFile("input.txt")

	squaresLeftInRow := map[Row]int{}
	squaresLeftInCol := map[Col]int{}

	// Start with empty rows and cols. Need 5 markings to continue
	numberOfBoards := len(inputs) / 6
	boardSums := map[int]int{}
	for board := 0; board < numberOfBoards; board++ {
		for i := 0; i < 5; i++ {
			squaresLeftInRow[Row{board, i}] = 5
			squaresLeftInCol[Col{board, i}] = 5
		}
		boardSums[board] = 0
	}

	// Set the locations of each number that gets called
	rowsOfChosenNumber := map[string]([]Row){}
	colsOfChosenNumber := map[string]([]Col){}
	i := 2
	board := 0
	for i < len(inputs) {
		for r := 0; r < 5; r++ {
			// TIL: Srings.spit doesn't work here!
			row := strings.Fields(inputs[i])
			for c := 0; c < 5; c++ {
				n := row[c]
				rowsOfChosenNumber[n] = append(rowsOfChosenNumber[n], Row{board, r})
				colsOfChosenNumber[n] = append(colsOfChosenNumber[n], Col{board, c})

				boardSums[board] += lib.ToInt(n)

			}
			i++
		}
		i++
		board++
	}

	// Let's play!!
	numbers := strings.Split(inputs[0], ",")
	hasBingo := false
	for _, n := range numbers {
		amt, _ := strconv.Atoi(n)

		// Mark rows
		rows := rowsOfChosenNumber[n]
		for _, row := range rows {
			boardSums[row.board] -= amt

			squaresLeftInRow[row] -= 1

			if squaresLeftInRow[row] == 0 {
				hasBingo = true
				print(amt*boardSums[row.board], "\n")
				break
			}
		}

		// Mark cols
		cols := colsOfChosenNumber[n]
		for _, col := range cols {
			squaresLeftInCol[col] -= 1

			if squaresLeftInCol[col] == 0 {
				hasBingo = true
				print(amt*boardSums[col.board], "\n")
				break
			}
		}

		if hasBingo {
			break
		}
	}

}

func PartTwo() {
	inputs := lib.ReadInputFile("input.txt")

	rowCheck := map[Row]int{}
	colCheck := map[Col]int{}

	// Each board's row / column has 5 left
	nBoards := len(inputs) / 6
	boardSums := map[int]int{}
	boardHasBingoed := map[int]bool{}
	for board := 0; board < nBoards; board++ {
		for i := 0; i < 5; i++ {
			rowCheck[Row{board, i}] = 5
			colCheck[Col{board, i}] = 5
		}
		boardSums[board] = 0
	}

	// Set the locations of each number that gets called
	rowsOfNumber := map[string]([]Row){}
	colsOfNumber := map[string]([]Col){}
	i := 2
	board := 0
	for i < len(inputs) {
		for r := 0; r < 5; r++ {
			row := strings.Fields(inputs[i])
			for c := 0; c < 5; c++ {
				n := row[c]
				rowsOfNumber[n] = append(rowsOfNumber[n], Row{board, r})
				colsOfNumber[n] = append(colsOfNumber[n], Col{board, c})

				amt, _ := strconv.Atoi(n)
				boardSums[board] += amt

			}
			i++
		}
		i++
		board++
	}

	// Let's play!!
	numbers := strings.Split(inputs[0], ",")
	sumFromLastBingo := 0
	for _, n := range numbers {
		amt := lib.ToInt(n)

		// Mark rows
		rows := rowsOfNumber[n]
		for _, row := range rows {
			boardSums[row.board] -= amt
			rowCheck[row] -= 1

			if rowCheck[row] == 0 && !boardHasBingoed[row.board] {
				boardHasBingoed[row.board] = true
				sumFromLastBingo = amt * boardSums[row.board]
			}
		}

		// Mark cols
		cols := colsOfNumber[n]
		for _, col := range cols {
			colCheck[col] -= 1
			if colCheck[col] == 0 && !boardHasBingoed[col.board] {
				boardHasBingoed[col.board] = true
				sumFromLastBingo = amt * boardSums[col.board]
			}

		}

	}
	print(sumFromLastBingo)
}

func main() {
	PartOne()

	PartTwo()
}
