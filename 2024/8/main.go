package main

import (
	"fmt"
	"io"
	"os"
)

type Node struct {
	X    int
	Y    int
	Type byte
}

func parseData(data []byte) (result map[byte][]Node, width int, height int) {

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

	result = make(map[byte][]Node, 0)

	for y := 0; y < len(lines); y++ {

		for x := 0; x < len(lines[0]); x++ {

			if lines[y][x] == '.' {
				continue
			}

			if _, exists := result[lines[y][x]]; !exists {
				result[lines[y][x]] = make([]Node, 0)
			}

			result[lines[y][x]] = append(
				result[lines[y][x]],
				Node{
					X:    x,
					Y:    y,
					Type: lines[y][x],
				},
			)

		}

	}

	width = len(lines[0])
	height = len(lines)

	return

}

func containsPos(positions [][2]int, pos [2]int) (result bool) {

	for _, v := range positions {
		result = v[0] == pos[0] && v[1] == pos[1]
		if result {
			return
		}
	}

	return

}

func partOne(nodeMap map[byte][]Node, width int, height int) (total int) {

	antiNodePositions := make([][2]int, 0)

	for _, nodes := range nodeMap {

		for i := 0; i < len(nodes); i++ {

			for j := 0; j < len(nodes); j++ {
				if nodes[i].X == nodes[j].X && nodes[i].Y == nodes[j].Y {
					continue
				}

				diff := [2]int{
					nodes[i].X - nodes[j].X,
					nodes[i].Y - nodes[j].Y,
				}

				antiNodePos := [2]int{
					nodes[i].X + diff[0],
					nodes[i].Y + diff[1],
				}

				if antiNodePos[0] >= 0 && antiNodePos[0] < width && antiNodePos[1] >= 0 && antiNodePos[1] < height {
					if !containsPos(antiNodePositions, antiNodePos) {
						antiNodePositions = append(antiNodePositions, antiNodePos)
					}
				}
			}

		}

	}

	return len(antiNodePositions)

}

func partTwo(nodeMap map[byte][]Node, width int, height int) (total int) {

	antiNodePositions := make([][2]int, 0)

	for _, nodes := range nodeMap {

		for i := 0; i < len(nodes); i++ {

			for j := 0; j < len(nodes); j++ {
				if nodes[i].X == nodes[j].X && nodes[i].Y == nodes[j].Y {
					continue
				}

				diff := [2]int{
					nodes[i].X - nodes[j].X,
					nodes[i].Y - nodes[j].Y,
				}

				var antiNodePos [2]int = [2]int{nodes[i].X, nodes[i].Y}
				if !containsPos(antiNodePositions, antiNodePos) {
					antiNodePositions = append(antiNodePositions, antiNodePos)
				}

				var inBounds bool = true

				for inBounds {
					antiNodePos = [2]int{
						antiNodePos[0] + diff[0],
						antiNodePos[1] + diff[1],
					}

					if antiNodePos[0] >= 0 && antiNodePos[0] < width && antiNodePos[1] >= 0 && antiNodePos[1] < height {
						if !containsPos(antiNodePositions, antiNodePos) {
							antiNodePositions = append(antiNodePositions, antiNodePos)
						}
					} else {
						inBounds = false
					}
				}

			}

		}

	}

	return len(antiNodePositions)

}

func main() {

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	parsed, width, height := parseData(data)

	fmt.Println(partOne(parsed, width, height))
	fmt.Println(partTwo(parsed, width, height))

}
