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

func containsPosition(visitedPositions [][3]int, pos [3]int) (idx int) {

	for i := len(visitedPositions) - 1; i > 0; i-- {

		if visitedPositions[i][0] == pos[0] && visitedPositions[i][1] == pos[1] && visitedPositions[i][2] == pos[2] {
			return i
		}

	}

	return -1

}

func partTwo(mapRows [][]byte, guardPosX int, guardPosY int) (totalLoops int) {

	guardStartX := guardPosX
	guardStartY := guardPosY

	for y := 0; y < len(mapRows); y++ {
		for x := 0; x < len(mapRows); x++ {

			guardPosX = guardStartX
			guardPosY = guardStartY

			if mapRows[y][x] == '#' || (guardPosX == x && guardPosY == y) {
				continue
			}

			mapRows[y][x] = obstruction

			var direction guardDirection = north
			var visitedPositions [][3]int = make([][3]int, 0)
			var possiblyInLoop bool
			var possibleLoopLength int
			var loopStartPos [3]int

			for guardPosX < len(mapRows[0])-1 && guardPosX > 0 && guardPosY < len(mapRows)-1 && guardPosY > 0 {

				contP := containsPosition(visitedPositions, [3]int{guardPosX, guardPosY, int(direction)})
				fmt.Printf("\r Obstruction XY: %d,%d...", x, y)

				if possiblyInLoop {

					if len(visitedPositions)-contP+1 == possibleLoopLength && loopStartPos[0] == guardPosX && loopStartPos[1] == guardPosY && loopStartPos[2] == int(direction) {
						totalLoops++
						break
					}

				} else if !possiblyInLoop && contP != -1 {
					possiblyInLoop = true
					possibleLoopLength = len(visitedPositions) - contP + 1
					loopStartPos = [3]int{guardPosX, guardPosY, int(direction)}
				}

				visitedPositions = append(visitedPositions, [3]int{guardPosX, guardPosY, int(direction)})

				// "Clamp" direction

				if direction > west {
					direction = north
				}

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

			mapRows[y][x] = emptySpace

		}
	}

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
