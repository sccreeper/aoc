package main

import (
	"fmt"
	"os"
	"strings"
	"time"
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

type Node struct {
	Data  string
	Left  *Node
	Right *Node
}

func make_tree(nodes map[string][2]string) map[string]*Node {

	//Convert the map of nodes into an array of node structs
	//This is my first ever tree IDK what i'm doing

	tree := make(map[string]*Node, 0)

	for k := range nodes {
		tree[k] = &Node{Data: k}
	}

	for k, v := range tree {

		tree[k].Left = tree[nodes[v.Data][0]]
		tree[k].Right = tree[nodes[v.Data][1]]

	}

	return tree

}

func part_1(directions string, nodes map[string][2]string) int {

	tree := make_tree(nodes)

	var z_found bool = false
	var current_node *Node = tree["AAA"]
	var direction_count int = 0

	for !z_found {

		for _, v := range directions {

			if current_node.Data == "ZZZ" {
				z_found = true
				break
			}

			if v == 'L' {
				current_node = current_node.Left
			} else {
				current_node = current_node.Right
			}

			direction_count++
		}

	}

	return direction_count

}

func part_2(directions string, nodes map[string][2]string) int {

	tree := make_tree(nodes)

	// Get starting nodes

	starting_nodes := make([]*Node, 0)

	for _, v := range tree {

		if strings.HasSuffix(v.Data, "A") {
			starting_nodes = append(starting_nodes, v)
		}

	}

	// Find number of steps for Z for each node

	var steps []int = make([]int, 0)

	for i, _ := range starting_nodes {

		var direction_count int = 0
		var z_found bool = false

		for !z_found {

			for _, d := range directions {

				if strings.HasSuffix(starting_nodes[i].Data, "Z") {
					z_found = true
					break
				}

				if d == 'L' {
					starting_nodes[i] = starting_nodes[i].Left
				} else {
					starting_nodes[i] = starting_nodes[i].Right
				}

				direction_count++

			}

		}

		steps = append(steps, direction_count)

	}

	return lcm(steps...)

}

func gcd(a int, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}

	return a
}

func lcm(integers ...int) int {

	result := integers[0] * integers[1] / gcd(integers[0], integers[1])

	for i := 2; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result

}

func main() {

	var file_name string
	fmt.Println("File name:")
	fmt.Scanln(&file_name)

	file_bytes, err := os.ReadFile(file_name)
	if err != nil {
		panic(err)
	}

	//fmt.Println(part_1(parse_file(file_bytes)))

	directions, nodes := parse_file(file_bytes)

	start := time.Now()

	fmt.Println(part_1(directions, nodes))

	fmt.Println(part_2(directions, nodes))

	fmt.Println(time.Since(start))

}
