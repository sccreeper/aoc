package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Game struct {
	ID       int
	Handfuls []Handful
}

type Handful struct {
	Red   int
	Green int
	Blue  int
}

const (
	RedLimit   int = 12
	GreenLimit int = 13
	BlueLimit  int = 14
)

// Shared
func main() {

	var file_name string
	fmt.Println("File name:")
	fmt.Scanln(&file_name)

	games_data, err := os.ReadFile(file_name)
	if err != nil {
		fmt.Printf("Error reading file %s", file_name)
	}

	start := time.Now()

	var lines []string = strings.Split(string(games_data), "\n")
	var games []Game = make([]Game, len(lines))

	// Parse file

	for i, line := range lines {

		games[i] = Game{
			ID:       i,
			Handfuls: make([]Handful, 0),
		}

		data := strings.TrimSpace(strings.Split(line, ":")[1])
		handfuls := strings.Split(data, ";")

		for j, handful := range handfuls {

			games[i].Handfuls = append(games[i].Handfuls, Handful{})

			cubes := strings.Split(handful, ",")

			for _, cube := range cubes {

				amount := strings.Split(strings.TrimSpace(cube), " ")

				x, _ := strconv.Atoi(amount[0])

				if amount[1] == "red" {

					games[i].Handfuls[j].Red = x

				} else if amount[1] == "green" {

					games[i].Handfuls[j].Green = x

				} else if amount[1] == "blue" {

					games[i].Handfuls[j].Blue = x

				}

			}

		}

	}

	// fmt.Println("Valid games are:")

	// var sum int

	// for _, game := range valid_games {

	// 	fmt.Printf("Game %d\n", game+1)
	// 	sum += game + 1

	// }

	fmt.Println(sum_of_powers(games))

	fmt.Printf("Used: %s\n", time.Since(start))

}

// Part 1
func get_valid_games(games []Game) []int {

	var valid_games []int = make([]int, 0)

	for i, game := range games {

		var game_valid bool = true

		for _, handful := range game.Handfuls {

			if handful.Red > RedLimit || handful.Green > GreenLimit || handful.Blue > BlueLimit {
				game_valid = false
				break
			}

		}

		if game_valid {
			valid_games = append(valid_games, i)
		}

	}

	return valid_games

}

// Part 2
func sum_of_powers(games []Game) int {

	var sum int = 0

	for _, game := range games {

		var red_min int = 1
		var green_min int = 1
		var blue_min int = 1

		for _, handful := range game.Handfuls {

			if handful.Red > red_min {
				red_min = handful.Red
			}

			if handful.Green > green_min {
				green_min = handful.Green
			}

			if handful.Blue > blue_min {
				blue_min = handful.Blue
			}

		}

		sum += red_min * green_min * blue_min

	}

	return sum

}
