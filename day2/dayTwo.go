package dayTwo

import (
	"aoc2023/utils"
	"fmt"
	"strconv"
	"strings"
)

type Game struct {
	Id     int
	Rounds []Round
}

func (g Game) IsValid() bool {
	isValid := true
	for _, r := range g.Rounds {
		if !r.IsValid() {
			isValid = false
			break
		}
	}
	return isValid
}

type Round struct {
	Red   int
	Green int
	Blue  int
}

func (r Round) IsValid() bool {
	if r.Red <= 12 && r.Green <= 13 && r.Blue <= 14 {
		return true
	}
	return false
}

var FileName string = "./day2/day2.txt"

func Run() {
	fmt.Println("Day 2 =========")
	fmt.Println("Part 1: ", PartOne())
}

func PartOne() int {
	fileScanner := utils.GetFileScanner(FileName)
	result := 0
	for fileScanner.Scan() {
		game := getGameDataFromString(fileScanner.Text())
		if game.IsValid() {
			result += game.Id
		}
	}

	return result
}

func PartTwo(line string) {

}

func getGameDataFromString(line string) Game {
	game := Game{}
	var roundData []string
	game.Id, roundData = getGameId(line)

	for _, r := range roundData {
		gameRound := Round{}
		colorsData := []string{}
		colorsData = append(colorsData, strings.Split(r, ",")...)
		for _, c := range colorsData {
			x := strings.Split(strings.TrimSpace(c), " ")
			colorType := string(x[1])
			n, err := strconv.Atoi(string(x[0]))
			if err != nil {
				fmt.Println(err)
			}

			switch colorType {
			case "red":
				gameRound.Red = n
			case "green":
				gameRound.Green = n
			case "blue":
				gameRound.Blue = n
			}
			game.Rounds = append(game.Rounds, gameRound)
		}
	}

	return game
}

func getGameId(input string) (int, []string) {
	data := strings.Split(input, ":")
	gameData := data[0]
	gameIdString := strings.Split(gameData, " ")
	gameId, err := strconv.Atoi(string(gameIdString[1]))
	if err != nil {
		fmt.Println(err)
	}

	gameRoundsData := data[1]
	roundData := strings.Split(gameRoundsData, ";")

	return gameId, roundData
}
