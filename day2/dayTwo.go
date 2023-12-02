package dayTwo

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type game struct {
	Id    int
	Red   int
	Green int
	Blue  int
}

func (g game) IsValid() bool {
	if g.Red <= 12 && g.Green <= 13 && g.Blue <= 14 {
		return true
	}
	return false
}

var FileName string = "./day2/day2.txt"

func Run() {
	fmt.Println("day two")
	fmt.Println("part one: ", PartOne())
}

func PartOne() int {
	fileScanner := getFileScanner()
	//games := []game{}
	result := 0
	for fileScanner.Scan() {
		g := game{}
		line := fileScanner.Text()
		data := strings.Split(line, ":")
		gameData := data[0]
		gameIdString := strings.Split(gameData, " ")
		gameId, err := strconv.Atoi(string(gameIdString[1]))
		if err != nil {
			fmt.Println(err)
		}
		g.Id = gameId
		gameRoundsData := data[1]
		roundData := strings.Split(gameRoundsData, ";")

		valid := make([]bool, len(roundData))
		for i := 0; i < len(valid); i++ {
			valid[i] = true
		}

		for idx, r := range roundData {
			colorsData := []string{}
			colorsData = append(colorsData, strings.Split(r, ",")...)
			for _, c := range colorsData {
				x := strings.Split(strings.TrimSpace(c), " ")
				colorType := string(x[1])
				n, err := strconv.Atoi(string(x[0]))
				if err != nil {
					fmt.Println(50, err)
				}
				switch colorType {
				case "red":
					g.Red = n
				case "green":
					g.Green = n
				case "blue":
					g.Blue = n
				}
			}
			valid[idx] = g.IsValid()
		}

		isValid := true
		for _, v := range valid {
			if !v {
				isValid = false
				break
			}
		}

		if isValid {
			result += g.Id
		}

	}

	return result
}

func PartTwo() {

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
