package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse_file(data []byte) []string {

	return strings.Split(string(data), ",")

}

func hash(data []byte) (result int) {
	for _, b := range data {

		result += int(b)
		result *= 17
		result %= 256

	}

	return
}

type Lens struct {
	Label    string
	Strength int
}

func contains_lens(data []Lens, label string) (contains bool, index int) {

	for i, v := range data {

		if v.Label == label {

			contains = true
			index = i
			return

		}

	}

	contains = false
	index = 0
	return

}

func remove_lens(lenses []Lens, label string) []Lens {

	var index int
	var exists bool = false

	for i, v := range lenses {

		if v.Label == label {
			index = i
			exists = true
			break
		}

	}

	if !exists {
		return lenses
	} else {
		return append(lenses[:index], lenses[index+1:]...)
	}

}

func part_2(data []string) (sum int) {

	boxes := make([][]Lens, 256)

	for i := 0; i < 256; i++ {
		boxes = append(boxes, make([]Lens, 0))
	}

	var location int = 0

	for _, instruction := range data {

		var label string
		var focal int

		if strings.Contains(instruction, "-") {
			label = strings.Split(instruction, "-")[0]
			location = hash([]byte(label))

			boxes[location] = remove_lens(boxes[location], label)

		} else {
			label = strings.Split(instruction, "=")[0]
			location = hash([]byte(label))
			focal, _ = strconv.Atoi(strings.Split(instruction, "=")[1])

			if exists, index := contains_lens(boxes[location], label); exists {
				boxes[location][index].Strength = focal
			} else {
				boxes[location] = append(boxes[location], Lens{Label: label, Strength: focal})
			}

		}

	}

	for i, box := range boxes {

		for j, l := range box {

			sum += (i + 1) * (j + 1) * l.Strength

		}

	}

	return

}

func part_1(data []string) (sum int) {

	sum = 0

	for _, v := range data {

		sum += hash([]byte(v))

	}

	return

}

func main() {

	var file_name string
	fmt.Println("File name:")
	fmt.Scanln(&file_name)

	file_data, err := os.ReadFile(file_name)
	if err != nil {
		panic(err)
	}

	fmt.Println(part_1(parse_file(file_data)))

	fmt.Println(part_2(parse_file(file_data)))

}
