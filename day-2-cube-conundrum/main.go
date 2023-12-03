package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	id       int
	draws    []Draw
	possible bool
}

type Draw struct {
	red   int
	blue  int
	green int
}

func parseGames(input string) []Game {
	var games []Game

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		parts := strings.Split(line, ":")
		gameNumber, _ := strconv.Atoi(strings.Trim(parts[0], "Game "))
		gameData := strings.Split(parts[1], ";")

		game := Game{id: gameNumber}
		for _, drawData := range gameData {
			draw := Draw{red: 0, blue: 0, green: 0}
			cubeCounts := strings.Split(strings.TrimSpace(drawData), ", ")
			for _, cubeCount := range cubeCounts {
				matches := strings.Fields(cubeCount)
				number, _ := strconv.Atoi(matches[0])
				color := matches[1]
				switch color {
				case "red":
					draw.red += number
				case "blue":
					draw.blue += number
				case "green":
					draw.green += number
				}
			}
			game.draws = append(game.draws, draw)
		}
		games = append(games, game)
	}

	return games
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func gameLimitCheck(game Game) bool {
	for _, draw := range game.draws {
		if draw.red > 12 || draw.blue > 14 || draw.green > 13 {
			return false
		}
	}
	return true
}

func calculateGamePower(game Game) int {
	// top seen color
	color := make(map[string]int)
	color["red"] = 0
	color["blue"] = 0
	color["green"] = 0
	for _, draw := range game.draws {
		color["red"] = int(math.Max(float64(color["red"]), float64(draw.red)))
		color["blue"] = int(math.Max(float64(color["blue"]), float64(draw.blue)))
		color["green"] = int(math.Max(float64(color["green"]), float64(draw.green)))
	}
	return color["red"] * color["blue"] * color["green"]
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("You need to provide an input file")
		return
	}

	fileInput := os.Args[1]

	fmt.Println("Input file: ", fileInput)

	rawData, err := os.ReadFile(fileInput)

	check(err)

	data := string(rawData)

	games := parseGames(data)

	// Part juan
	sumGamesPossible := 0
	for _, game := range games {
		if gameLimitCheck(game) {
			fmt.Println("Game ", game.id, " is possible")
			sumGamesPossible += game.id
		} else {
			fmt.Println("Game ", game.id, " is not possible")
		}
	}
	fmt.Println("Sum of all possible games: ", sumGamesPossible)
	// Part dos
	total := 0
	for _, game := range games {
		total += calculateGamePower(game)
	}
	fmt.Println("Total power of all games: ", total)
}
