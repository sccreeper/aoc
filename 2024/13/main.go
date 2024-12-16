package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

type game struct {
	ButtonA [2]int
	ButtonB [2]int
	Prize   [2]int
}

func parseData(data []byte) (result []game) {

	result = make([]game, 0)

	lines := make([]string, 0)

	accumulator := ""

	for _, b := range data {

		if b == '\n' {
			lines = append(lines, accumulator)
			accumulator = ""
		} else {
			accumulator += string(b)
		}

	}

	lines = append(lines, accumulator)

	for i := 0; i < len(lines); i += 4 {

		currentGame := game{}

		buttonALine := strings.Split(lines[i], ":")[1]
		buttonBLine := strings.Split(lines[i+1], ":")[1]
		prizeLine := strings.Split(lines[i+2], ":")[1]

		button := strings.Split(buttonALine, ",")
		buttonX := strings.Replace(button[0], " X+", "", 1)
		buttonY := strings.Replace(button[1], " Y+", "", 1)

		buttonXNum, _ := strconv.Atoi(buttonX)
		buttonYNum, _ := strconv.Atoi(buttonY)

		currentGame.ButtonA = [2]int{buttonXNum, buttonYNum}

		button = strings.Split(buttonBLine, ",")
		buttonX = strings.Replace(button[0], " X+", "", 1)
		buttonY = strings.Replace(button[1], " Y+", "", 1)

		buttonXNum, _ = strconv.Atoi(buttonX)
		buttonYNum, _ = strconv.Atoi(buttonY)

		currentGame.ButtonB = [2]int{buttonXNum, buttonYNum}

		prize := strings.Split(prizeLine, ",")
		buttonX = strings.Replace(prize[0], " X=", "", 1)
		buttonY = strings.Replace(prize[1], " Y=", "", 1)

		buttonXNum, _ = strconv.Atoi(buttonX)
		buttonYNum, _ = strconv.Atoi(buttonY)

		currentGame.Prize = [2]int{buttonXNum, buttonYNum}

		result = append(result, currentGame)

	}

	return

}

func partOne(games []game) (numTokens int) {

	for _, v := range games {

		var aPresses float64
		var bPresses float64

		aPresses = float64((v.Prize[0]*v.ButtonB[1])-(v.Prize[1]*v.ButtonB[0])) / float64((v.ButtonA[0]*v.ButtonB[1])-(v.ButtonA[1]*v.ButtonB[0]))
		bPresses = float64((v.Prize[1]*v.ButtonA[0])-(v.Prize[0]*v.ButtonA[1])) / float64((v.ButtonA[0]*v.ButtonB[1])-(v.ButtonA[1]*v.ButtonB[0]))

		aCeil := math.Ceil(aPresses)
		bCeil := math.Ceil(bPresses)

		if aPresses != aCeil || bPresses != bCeil {
			continue
		} else if aPresses > 100 || bPresses > 100 {
			continue
		} else {
			numTokens += (int(aPresses) * 3) + int(bPresses)
		}

	}

	return

}

func partTwo(games []game) (numTokens int) {

	for i := 0; i < len(games); i++ {

		games[i] = game{
			ButtonA: games[i].ButtonA,
			ButtonB: games[i].ButtonB,
			Prize: [2]int{
				games[i].Prize[0] + 10000000000000,
				games[i].Prize[1] + 10000000000000,
			},
		}

	}

	for _, v := range games {

		var aPresses float64
		var bPresses float64

		aPresses = float64((v.Prize[0]*v.ButtonB[1])-(v.Prize[1]*v.ButtonB[0])) / float64((v.ButtonA[0]*v.ButtonB[1])-(v.ButtonA[1]*v.ButtonB[0]))
		bPresses = float64((v.Prize[1]*v.ButtonA[0])-(v.Prize[0]*v.ButtonA[1])) / float64((v.ButtonA[0]*v.ButtonB[1])-(v.ButtonA[1]*v.ButtonB[0]))

		aCeil := math.Ceil(aPresses)
		bCeil := math.Ceil(bPresses)

		if aPresses != aCeil || bPresses != bCeil {
			continue
		} else {
			numTokens += (int(aPresses) * 3) + int(bPresses)
		}

	}

	return

}

func main() {

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	parsed := parseData(data)

	fmt.Println(partOne(parsed))
	fmt.Println(partTwo(parsed))

}
