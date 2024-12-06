package main

import (
	"fmt"
	"io"
	"os"
)

const emptySpace = '.'
const visited = 'X'
const obstruction = '#'
const guard = '^'

type guardDirection byte

const (
	north guardDirection = 0
	east  guardDirection = 1
	south guardDirection = 2
	west  guardDirection = 3
)

func parseData(data []byte) (mapRows [][]byte, guardPosX int, guardPosY int) {

	accumulator := make([]byte, 0)
	mapRows = make([][]byte, 0)

	for _, v := range data {

		if v == '\n' {

			mapRows = append(mapRows, accumulator)
			accumulator = make([]byte, 0)

		} else {

			if v == guard {
				accumulator = append(accumulator, emptySpace)

				guardPosY = len(mapRows)
				guardPosX = len(accumulator) - 1

			} else {
				accumulator = append(accumulator, v)
			}

		}

	}

	mapRows = append(mapRows, accumulator)

	return

}

func partOne(mapRows [][]byte, guardPosX int, guardPosY int) (totalVisited int) {

	var direction guardDirection = north

	for guardPosX < len(mapRows[0])-1 && guardPosX > 0 && guardPosY < len(mapRows)-1 && guardPosY > 0 {

		// "Clamp" direction

		if direction > west {
			direction = north
		}

		mapRows[guardPosY][guardPosX] = visited

		var moveDirection [2]int

		switch direction {
		case north:
			moveDirection = [2]int{0, -1}
		case east:
			moveDirection = [2]int{1, 0}
		case south:
			moveDirection = [2]int{0, 1}
		case west:
			moveDirection = [2]int{-1, 0}
		default:
			panic("wrong direction")
		}

		if mapRows[guardPosY+moveDirection[1]][guardPosX+moveDirection[0]] == obstruction {
			direction++
		} else {

			guardPosX += moveDirection[0]
			guardPosY += moveDirection[1]

		}

	}

	// Scan through visited cells and add to total

	for i := 0; i < len(mapRows); i++ {
		for j := 0; j < len(mapRows[0]); j++ {

			if mapRows[i][j] == visited {
				totalVisited++
			}

		}
	}

	// Account for the cell when the guard left the grid
	totalVisited++

	return

}

// I will try this another time
// Tried like 3 different implementations and none of them worked.

func checkForLoop() {

}

func partTwo(mapRows [][]byte, guardPosX int, guardPosY int) (totalLoops int) {
	return
}

func main() {

	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	mapRows, guardPosX, guardPosY := parseData(bytes)
	fmt.Printf("Guard X: %d Guard Y: %d\n", guardPosX, guardPosY)

	fmt.Println(partOne(mapRows, guardPosX, guardPosY))
	fmt.Printf("\nTotal loops: %d\n", partTwo(mapRows, guardPosX, guardPosY))

}
