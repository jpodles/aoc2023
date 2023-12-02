package dayOne

import (
	"aoc2023/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var FileName string = "./day1/day1.txt"

var numbersMap map[string]string = map[string]string{
	"one":   "o1e",
	"two":   "t2o",
	"three": "t3e",
	"four":  "f4r",
	"five":  "f5e",
	"six":   "s6x",
	"seven": "s7n",
	"eight": "e8t",
	"nine":  "n9e",
}

func Run() {
	fmt.Println("Day 1 =========")
	fmt.Println("Part 1: ", partOne())
	fmt.Println("Part 2: ", partTwo())
}

func partOne() int {
	fileScanner := utils.GetFileScanner(FileName)
	result := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		result += getNumber(line)
	}
	return result
}

func partTwo() int {
	fileScanner := utils.GetFileScanner(FileName)
	result := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		for key, value := range numbersMap {
			if strings.Contains(line, key) {
				line = strings.ReplaceAll(line, key, value)
			}
		}
		result += getNumber(line)
	}
	return result
}

func getNumber(line string) int {
	var first, last string
	runes := []rune(line)
	for _, r := range runes {
		if unicode.IsDigit(r) {
			first = string(r)
			break
		}
	}

	for i := len(runes) - 1; i >= 0; i-- {
		r := runes[i]
		if unicode.IsDigit(r) {
			last = string(r)
			break
		}
	}

	number, err := strconv.Atoi(first + last)
	if err != nil {
		fmt.Println(err)
	}
	return number
}

func getFileScanner() *bufio.Scanner {
	readFile, err := os.Open(FileName)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	return fileScanner
}
