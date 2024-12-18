package main

import (
	"fmt"
	"io"
	"os"
)

type mazeNode struct {
	Character byte
	Adjacent  []*mazeNode
	Position  [2]int
}

var lookVectors [][2]int = [][2]int{
	{0, -1},
	{0, 1},
	{1, 0},
	{-1, 0},
}

type reindeerDirection uint8

const (
	north reindeerDirection = 0
	east  reindeerDirection = 1
	south reindeerDirection = 2
	west  reindeerDirection = 3
)

var directionVectorsMap map[[2]int]reindeerDirection = map[[2]int]reindeerDirection{
	lookVectors[0]: north,
	lookVectors[1]: south,
	lookVectors[2]: east,
	lookVectors[3]: west,
}

const wallChar byte = '#'
const freeSpaceChar byte = '.'
const startChar byte = 'S'
const endChar byte = 'E'

func containsNode(nodes []*mazeNode, node *mazeNode) bool {

	for _, n := range nodes {

		if n.Position == node.Position && n.Character == node.Character {
			return true
		}

	}

	return false

}

func parseData(data []byte) (start *mazeNode) {

	accumulator := make([]byte, 0)
	lines := make([][]byte, 0)

	for _, v := range data {

		if v == '\n' {
			lines = append(lines, accumulator)
			accumulator = make([]byte, 0)
		} else {
			accumulator = append(accumulator, v)
		}

	}

	lines = append(lines, accumulator)

	nodes := make([][]*mazeNode, len(lines))

	// Begin graph

	var startNodePos [2]int

	for y := 0; y < len(lines); y++ {
		nodes[y] = make([]*mazeNode, len(lines[0]))

		for x := 0; x < len(lines[0]); x++ {
			nodes[y][x] = &mazeNode{
				Character: lines[y][x],
				Position:  [2]int{x, y},
				Adjacent:  make([]*mazeNode, 0),
			}

			if lines[y][x] == startChar {
				startNodePos = [2]int{x, y}
			}
		}

	}

	for y := 0; y < len(nodes); y++ {
		for x := 0; x < len(nodes[0]); x++ {

			if nodes[y][x].Character == wallChar {
				continue
			} else {

				for _, vec := range lookVectors {

					if x+vec[0] < 0 || y+vec[0] < 0 || x+vec[0] >= len(lines[0]) || y+vec[0] >= len(lines[0]) {
						continue
					} else {

						if !containsNode(nodes[y+vec[1]][x+vec[0]].Adjacent, nodes[y][x]) {
							nodes[y][x].Adjacent = append(nodes[y][x].Adjacent, nodes[y+vec[1]][x+vec[0]])
						}

					}

				}

			}

		}
	}

	return nodes[startNodePos[1]][startNodePos[0]]

}

func walk(start *mazeNode, direction reindeerDirection, exploredNodes []*mazeNode) (foundNodes []*mazeNode, foundEnd bool, score int) {

	return

}

func partOne(start *mazeNode) (score int) {

	_, _, score = walk(start, east, make([]*mazeNode, 0))
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
