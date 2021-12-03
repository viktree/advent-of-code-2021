package lib

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

func ReadInputFile(relativeFilePath string) []string {
	// convert to absolute path
	absPath, err := filepath.Abs(relativeFilePath)
	if err != nil {
		log.Fatal("Error: failed to get entire filepath.")
	}

	// read from file
	f, err := os.Open(absPath)
	if err != nil {
		log.Fatal("Error: failed to open file.")
		os.Exit(2)
	}
	defer f.Close()

	var lines []string
	fileScanner := bufio.NewScanner(f)
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	return lines
}

func ReadInputFileToIntegers(relativeFilePath string) []int {
	// convert to absolute path
	absPath, err := filepath.Abs(relativeFilePath)
	if err != nil {
		log.Fatal("Error: failed to get entire filepath.")
	}

	// read from file
	f, err := os.Open(absPath)

	if err != nil {
		log.Fatal("Error: failed to open file.")
		os.Exit(2)
	}
	defer f.Close()

	var lines []int
	fileScanner := bufio.NewScanner(f)
	for fileScanner.Scan() {
		n, _ := strconv.Atoi(fileScanner.Text())
		lines = append(lines, n)
	}

	return lines
}
