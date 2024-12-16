package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type robot struct {
	Pos      [2]int
	Velocity [2]int
}

const spaceWidth = 101
const spaceHeight = 103

func parseData(data []byte) (result map[[2]int][]robot) {

	result = make(map[[2]int][]robot)

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

	for _, v := range lines {

		posString := strings.Split(v, " ")[0]
		velString := strings.Split(v, " ")[1]

		posString = strings.ReplaceAll(posString, "p=", "")
		velString = strings.ReplaceAll(velString, "v=", "")

		posX, err := strconv.Atoi(strings.Split(posString, ",")[0])
		posY, err := strconv.Atoi(strings.Split(posString, ",")[1])

		velX, err := strconv.Atoi(strings.Split(velString, ",")[0])
		velY, err := strconv.Atoi(strings.Split(velString, ",")[1])

		if err != nil {
			panic(err)
		}

		result[[2]int{posX, posY}] = append(result[[2]int{posX, posY}], robot{
			Pos:      [2]int{posX, posY},
			Velocity: [2]int{velX, velY},
		})

	}

	return

}

func prettifyMap(robots map[[2]int][]robot) {

	for y := 0; y < spaceHeight; y++ {

		lenStr := ""

		for x := 0; x < spaceWidth; x++ {

			if len(robots[[2]int{x, y}]) != 0 {
				lenStr += strconv.Itoa(len(robots[[2]int{x, y}]))
			} else {
				lenStr += "."
			}

		}

		fmt.Println(lenStr)
	}

}

func partOne(robots map[[2]int][]robot) (result int) {

	for i := 0; i < 100; i++ {

		newRobots := make(map[[2]int][]robot, 0)

		for _, v := range robots {

			for _, r := range v {

				newPos := [2]int{
					r.Pos[0] + r.Velocity[0],
					r.Pos[1] + r.Velocity[1],
				}

				if newPos[0] < 0 {
					newPos[0] += spaceWidth
				}

				if newPos[0] >= spaceWidth {
					newPos[0] -= spaceWidth
				}

				if newPos[1] < 0 {
					newPos[1] += spaceHeight
				}

				if newPos[1] >= spaceHeight {
					newPos[1] -= spaceHeight
				}

				newRobots[newPos] = append(
					newRobots[newPos],
					robot{
						Pos:      newPos,
						Velocity: r.Velocity,
					},
				)

			}

		}

		robots = newRobots

	}

	quadWidth := (spaceWidth / 2) - 1
	quadHeight := (spaceHeight / 2) - 1

	quadTotals := [4]int{}

	for pos, v := range robots {

		if pos[0] == quadWidth+1 || pos[1] == quadHeight+1 || len(v) == 0 {
			continue
		} else {

			// Figure out quadrant

			quadIndex := 0

			if pos[0] <= quadWidth {
				if pos[1] <= quadHeight {
					quadIndex = 0
				} else {
					quadIndex = 2
				}
			} else {
				if pos[1] <= quadHeight {
					quadIndex = 1
				} else {
					quadIndex = 3
				}
			}

			quadTotals[quadIndex] += len(v)

		}

	}

	result = 1

	for _, v := range quadTotals {
		if v == 0 {
			continue
		}

		result *= v
	}

	//prettifyMap(robots)

	return

}

func main() {

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	parsed := parseData(data)

	fmt.Println(partOne(parsed))

}
