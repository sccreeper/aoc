package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

type Node struct {
	Value   int
	Visited bool
}

var lookVectors [][2]int = [][2]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func parseData(data []byte) (trailMap [][]Node, headPositions [][2]int) {

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

	// Parse numbers

	for y := 0; y < len(lines); y++ {
		trailMap = append(trailMap, make([]Node, len(lines[0])))

		for x := 0; x < len(lines[0]); x++ {

			num, err := strconv.Atoi(string(lines[y][x]))
			if err != nil {
				panic(err)
			}

			trailMap[y][x] = Node{num, false}

			if num == 0 {
				headPositions = append(headPositions, [2]int{x, y})
			}

		}
	}

	return
}

func outOfBounds(pos [2]int, width int, height int) bool {
	return pos[0] < 0 || pos[1] < 0 || pos[0] >= width || pos[1] >= height
}

func prettyifyTrailMap(trailMap [][]Node) (prettified string) {

	prettified = ""

	for _, row := range trailMap {
		line := ""

		for _, n := range row {
			if n.Visited {
				line += strconv.Itoa(n.Value)
			} else {
				line += "."
			}
		}

		prettified += fmt.Sprintf("\n%s", line)
	}

	return

}

func walk(trailMap [][]Node, position [2]int) (hits int) {

	var ableToContinue bool = true

	for ableToContinue {

		trailMap[position[1]][position[0]].Visited = true

		if trailMap[position[1]][position[0]].Value == 9 {
			return 1
		}

		ableToContinue = false

		possibleBranches := make([][2]int, 0)

		for _, v := range lookVectors {

			newPosition := [2]int{position[0] + v[0], position[1] + v[1]}

			if outOfBounds(newPosition, len(trailMap[0]), len(trailMap)) {
				continue
			}

			if trailMap[newPosition[1]][newPosition[0]].Value-trailMap[position[1]][position[0]].Value == 1 &&
				!trailMap[newPosition[1]][newPosition[0]].Visited {
				possibleBranches = append(possibleBranches, newPosition)
			}

		}

		if len(possibleBranches) == 0 {
			ableToContinue = false
			break
		} else if len(possibleBranches) == 1 {
			position = possibleBranches[0]
			ableToContinue = true
		} else {

			for _, b := range possibleBranches {

				newHits := walk(trailMap, b)
				if newHits >= 1 {
					hits += newHits
				}

			}

		}

	}

	return

}

func partOne(trailMap [][]Node, headPositions [][2]int) (total int) {

	for _, v := range headPositions {

		// Reset trail map visited positions
		for y := 0; y < len(trailMap); y++ {
			for x := 0; x < len(trailMap[0]); x++ {
				trailMap[y][x].Visited = false
			}
		}

		newHits := walk(trailMap, v)
		total += newHits

	}

	return
}

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	trailMap, headPositions := parseData(data)
	fmt.Println(partOne(trailMap, headPositions))
}
