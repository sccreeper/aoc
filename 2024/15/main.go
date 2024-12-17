package main

import (
	"fmt"
	"io"
	"os"
)

const directionUp byte = '^'
const directionDown byte = 'v'
const directionLeft byte = '<'
const directionRight byte = '>'

var directionVectors map[byte][2]int = map[byte][2]int{
	directionUp:    {0, -1},
	directionDown:  {0, 1},
	directionLeft:  {-1, 0},
	directionRight: {1, 0},
}

const obstruction byte = '#'
const robot byte = '@'
const box byte = 'O'
const expandedBoxLeft byte = '['
const expandedBoxRight byte = ']'
const emptySpace byte = '.'

func parseData(data []byte) (maze [][]byte, instructions []byte, guardPos [2]int) {

	accumulator := make([]byte, 0)
	var inMaze bool = true

	guardPos = [2]int{-1, -1}

	for _, v := range data {

		if v == '\n' {

			if len(accumulator) == 0 {
				inMaze = false
				continue
			} else {

				if inMaze {
					maze = append(maze, accumulator)
				} else {
					instructions = append(instructions, accumulator...)
				}

				accumulator = make([]byte, 0)

			}

		} else {

			accumulator = append(accumulator, v)

		}

	}

	instructions = append(instructions, accumulator...)

	for y := 0; y < len(maze); y++ {
		for x := 0; x < len(maze[0]); x++ {

			if maze[y][x] == robot {
				guardPos = [2]int{x, y}
				break
			}

		}

		if guardPos[0] != -1 && guardPos[1] != -1 {
			break
		}
	}

	return

}

func expandMap(maze [][]byte) (expanded [][]byte, guardPos [2]int) {

	expanded = make([][]byte, 0)

	for y := 0; y < len(maze); y++ {
		expanded = append(expanded, make([]byte, 0))

		for x := 0; x < len(maze[0]); x++ {

			switch maze[y][x] {
			case emptySpace:
				expanded[y] = append(expanded[y], ".."...)
			case obstruction:
				expanded[y] = append(expanded[y], "##"...)
			case box:
				expanded[y] = append(expanded[y], "[]"...)
			case robot:
				expanded[y] = append(expanded[y], "@."...)
			default:
				continue
			}

		}
	}

	guardPos = [2]int{-1, -1}

	for y := 0; y < len(maze); y++ {
		for x := 0; x < len(maze[0]); x++ {

			if expanded[y][x] == robot {
				guardPos = [2]int{x, y}
				break
			}

		}

		if guardPos[0] != -1 && guardPos[1] != -1 {
			break
		}
	}

	return

}

func prettifyMap(maze [][]byte) {

	for _, v := range maze {

		fmt.Println(string(v))

	}

}

func partOne(maze [][]byte, instructions []byte, robotPos [2]int) (result int) {

	// fmt.Println(string(instructions))

	for _, v := range instructions {

		lookDirection := directionVectors[v]
		newPos := [2]int{robotPos[0] + lookDirection[0], robotPos[1] + lookDirection[1]}
		// fmt.Println(newPos)
		// fmt.Println(string(v))
		// prettifyMap(maze)

		if newPos[0] <= 0 || newPos[1] <= 0 || newPos[0] >= len(maze[0])-1 || newPos[1] >= len(maze)-1 {
			continue
		} else if maze[newPos[1]][newPos[0]] == obstruction {
			continue
		}

		if maze[newPos[1]][newPos[0]] == box {

			var ableToMove bool

			lookPos := newPos

			for {

				if lookPos[0] <= 0 || lookPos[1] <= 0 || lookPos[0] >= len(maze[0])-1 || lookPos[1] >= len(maze)-1 {
					break
				} else if maze[lookPos[1]][lookPos[0]] == obstruction {
					break
				}

				if maze[lookPos[1]][lookPos[0]] == emptySpace {
					ableToMove = true
					break
				}

				lookPos = [2]int{lookPos[0] + lookDirection[0], lookPos[1] + lookDirection[1]}

			}

			if ableToMove {

				// Swap boxes around

				maze[robotPos[1]][robotPos[0]] = emptySpace
				maze[newPos[1]][newPos[0]] = robot
				maze[lookPos[1]][lookPos[0]] = box

				robotPos = newPos

			}

		} else if maze[newPos[1]][newPos[0]] == emptySpace {
			maze[robotPos[1]][robotPos[0]] = emptySpace
			maze[newPos[1]][newPos[0]] = robot

			robotPos = newPos
		}

	}

	for y := 0; y < len(maze); y++ {
		for x := 0; x < len(maze[0]); x++ {

			if maze[y][x] == box {
				result += (100 * y) + x
			}

		}
	}

	return

}

func partTwo(maze [][]byte, instructions []byte, robotPos [2]int) {

}

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	maze, instructions, guardPos := parseData(data)

	fmt.Println(partOne(maze, instructions, guardPos))
}
