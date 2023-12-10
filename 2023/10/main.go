package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
	"time"
)

type Pipe int

const (
	VerticalPipe   Pipe = 0
	HorizontalPipe Pipe = 1
	NorthEastBend  Pipe = 2
	NorthWestBend  Pipe = 3
	SouthWestBend  Pipe = 4
	SouthEastBend  Pipe = 5
	Ground         Pipe = 6
	Start          Pipe = 7
)

var pipe_chars map[rune]Pipe = map[rune]Pipe{
	'|': VerticalPipe,
	'-': HorizontalPipe,
	'L': NorthEastBend,
	'J': NorthWestBend,
	'7': SouthWestBend,
	'F': SouthEastBend,
	'.': Ground,
	'S': Start,
}

type PipeNode struct {
	Type        Pipe
	Location    [2]int
	Connections []*PipeNode
}

var possible_connections map[rune]map[Pipe][]Pipe = map[rune]map[Pipe][]Pipe{
	'N': { //Up
		HorizontalPipe: {},
		VerticalPipe:   {SouthEastBend, SouthWestBend, VerticalPipe, Start},
		NorthEastBend:  {VerticalPipe, SouthEastBend, SouthWestBend, Start},
		NorthWestBend:  {VerticalPipe, SouthEastBend, SouthWestBend, Start},
		SouthEastBend:  {},
		SouthWestBend:  {},
		Start:          {VerticalPipe, SouthEastBend, SouthWestBend},
	},
	'S': { //Down
		HorizontalPipe: {},
		VerticalPipe:   {NorthEastBend, NorthWestBend, VerticalPipe, Start},
		NorthEastBend:  {},
		NorthWestBend:  {},
		SouthEastBend:  {VerticalPipe, NorthEastBend, NorthWestBend, Start},
		SouthWestBend:  {VerticalPipe, NorthEastBend, NorthWestBend, Start},
		Start:          {VerticalPipe, NorthEastBend, NorthWestBend},
	},
	'E': { //Right
		HorizontalPipe: {NorthWestBend, SouthWestBend, HorizontalPipe, Start},
		VerticalPipe:   {},
		NorthEastBend:  {HorizontalPipe, SouthWestBend, NorthWestBend, Start},
		NorthWestBend:  {},
		SouthEastBend:  {HorizontalPipe, SouthWestBend, NorthWestBend, Start},
		SouthWestBend:  {},
		Start:          {HorizontalPipe, SouthWestBend, NorthWestBend},
	},
	'W': { //Left
		HorizontalPipe: {NorthEastBend, SouthEastBend, HorizontalPipe, Start},
		VerticalPipe:   {},
		NorthEastBend:  {},
		NorthWestBend:  {NorthEastBend, SouthEastBend, HorizontalPipe, Start},
		SouthEastBend:  {},
		SouthWestBend:  {NorthEastBend, SouthEastBend, HorizontalPipe, Start},
		Start:          {NorthEastBend, SouthEastBend, HorizontalPipe},
	},
}

var opposite_directions map[rune]rune = map[rune]rune{
	'N': 'S',
	'S': 'N',
	'E': 'W',
	'W': 'E',
}

func can_connect(a Pipe, b Pipe, direction rune) bool {

	return slices.Contains(possible_connections[direction][a], b) && slices.Contains(possible_connections[opposite_directions[direction]][b], a)

}

func parse_file(data []byte) ([2]int, [][]PipeNode) {

	data_string := strings.Split(string(data), "\n")
	pipes := make([][]Pipe, len(data))

	for y := 0; y < len(data_string); y++ {

		pipes[y] = make([]Pipe, 0)

		for x := 0; x < len(data_string[y]); x++ {

			pipes[y] = append(pipes[y], pipe_chars[rune(data_string[y][x])])

		}

	}

	nodes := make([][]PipeNode, len(pipes))

	//Create all initial nodes

	start_coords := [2]int{}

	for y := 0; y < len(pipes); y++ {

		nodes[y] = make([]PipeNode, len(pipes[0]))

		if len(pipes[y]) == 0 {
			continue
		}

		for x := 0; x < len(pipes[0]); x++ {

			nodes[y][x] = PipeNode{
				Type:     pipes[y][x],
				Location: [2]int{x, y},
			}

			if nodes[y][x].Type == Start {
				start_coords = [2]int{x, y}
			}

		}

	}

	//Set connected pipes for all nodes.

	for y := 0; y < len(nodes); y++ {

		for x := 0; x < len(nodes[0]); x++ {

			if nodes[y][x].Type == Ground {
				continue
			}

			nodes[y][x].Connections = make([]*PipeNode, 0)
			//Check north connections

			if y != 0 && can_connect(nodes[y][x].Type, nodes[y-1][x].Type, 'N') {
				nodes[y][x].Connections = append(nodes[y][x].Connections, &nodes[y-1][x])
			}

			//Check south connections

			if y != len(nodes)-1 && can_connect(nodes[y][x].Type, nodes[y+1][x].Type, 'S') {
				nodes[y][x].Connections = append(nodes[y][x].Connections, &nodes[y+1][x])
			}

			//Check west connections

			if x != 0 && can_connect(nodes[y][x].Type, nodes[y][x-1].Type, 'W') {
				nodes[y][x].Connections = append(nodes[y][x].Connections, &nodes[y][x-1])
			}

			//Check east connections

			if x != len(nodes[0])-1 && can_connect(nodes[y][x].Type, nodes[y][x+1].Type, 'E') {
				nodes[y][x].Connections = append(nodes[y][x].Connections, &nodes[y][x+1])
			}

		}

	}

	return start_coords, nodes

}

func part_1(start [2]int, nodes [][]PipeNode) int {

	var steps int = 0
	var current_node *PipeNode
	var previous_node *PipeNode
	var s_found = false

	previous_node = &nodes[start[1]][start[0]]

	if nodes[start[1]][start[0]].Connections[0].Type == Start {
		current_node = nodes[start[1]][start[0]].Connections[1]
	} else {
		current_node = nodes[start[1]][start[0]].Connections[0]
	}

	for !s_found {

		if current_node.Connections[0] == previous_node {
			previous_node = current_node
			current_node = previous_node.Connections[1]
		} else {
			previous_node = current_node
			current_node = previous_node.Connections[0]
		}

		steps++

		if current_node.Type == Start {
			s_found = true
		}

	}

	return (steps + 1) / 2

}

func main() {

	var file_name string
	fmt.Println("File name:")
	fmt.Scanln(&file_name)

	file_data, err := os.ReadFile(file_name)
	if err != nil {
		panic(err)
	}

	used := time.Now()

	start, nodes := parse_file(file_data)

	fmt.Println(part_1(start, nodes))

	fmt.Println(time.Since(used))

}
