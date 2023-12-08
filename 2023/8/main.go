package main

import (
	"fmt"
	"os"
	"strings"
)

func parse_file(data []byte) (directions string, nodes map[string][2]string) {

	lines := strings.Split(string(data), "\n")

	directions = lines[0]

	nodes = make(map[string][2]string)

	for _, v := range lines[2:] {

		node := v[:3]
		left_direction := v[7:10]
		right_direction := v[12:15]

		// fmt.Printf("Node %s\n", node)
		// fmt.Println(left_direction)
		// fmt.Println(right_direction)

		nodes[node] = [2]string{left_direction, right_direction}

	}

	return

}

func part_1(directions string, nodes map[string][2]string) int {

	var z_found bool = false
	var current_node string = "AAA"
	var direction_count int = 0

	for !z_found {

		for _, v := range directions {

			if current_node == "ZZZ" {
				z_found = true
				break
			}

			if v == 'L' {
				current_node = nodes[current_node][0]
			} else {
				current_node = nodes[current_node][1]
			}

			direction_count++
		}

	}

	return direction_count

}

func main() {

	var file_name string
	fmt.Println("File name:")
	fmt.Scanln(&file_name)

	file_bytes, err := os.ReadFile(file_name)
	if err != nil {
		panic(err)
	}

	fmt.Println(part_1(parse_file(file_bytes)))

}
