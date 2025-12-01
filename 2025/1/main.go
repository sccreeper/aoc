package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Rotation struct {
	Direction byte
	Distance  int
}

func parse(data []byte) (result []Rotation) {

	lines := strings.Split(string(data), "\n")

	result = make([]Rotation, 0)

	for _, v := range lines {

		if len(v) == 0 {
			continue
		}

		distance, err := strconv.Atoi(v[1:])
		if err != nil {
			panic(err)
		}

		result = append(result, Rotation{
			Direction: v[0],
			Distance:  distance,
		})

	}

	return

}

func part_one(data []Rotation) (result int) {

	var rotation_value int = 50

	for _, v := range data {

		if v.Direction == 'L' {
			rotation_value -= v.Distance
		} else {
			rotation_value += v.Distance
		}

		for rotation_value < 0 || rotation_value > 99 {
			if rotation_value < 0 {
				rotation_value += 100
			}

			if rotation_value > 99 {
				rotation_value -= 100
			}
		}

		if rotation_value == 0 {
			result++
		}

	}

	return

}

func part_two(data []Rotation) (result int) {

	var rotation_value int = 50

	for _, v := range data {

		rotations := v.Distance / 100
		result += rotations

		for i := 0; i < v.Distance%100; i++ {
			if v.Direction == 'L' {
				rotation_value--
			} else {
				rotation_value++
			}

			rotation_value %= 100

			if rotation_value == 0 {
				result++
			}
		}

	}

	return

}

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	rotations := parse(data)

	fmt.Println(part_one(rotations))
	fmt.Println(part_two(rotations))

}
