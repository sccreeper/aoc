package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

const numbers string = "1234567890"

type side int

// No diagonal: If Left and Top are present this means we know that the diagonal top left isn't out of bounds.
const (
	Left   side = 0
	Right  side = 1
	Top    side = 2
	Bottom side = 3
)

func main() {

	var file_name string

	fmt.Println("File name:")
	fmt.Scanln(&file_name)

	engine_bytes, err := os.ReadFile(file_name)
	if err != nil {
		fmt.Printf("Error reading file %s", file_name)
		return
	}

	start := time.Now()

	engine_lines := strings.Split(string(engine_bytes), "\n")

	gears := make([][][]int, len(engine_lines))
	for i := range gears {
		gears[i] = make([][]int, len(engine_lines[0]))
	}
	for i := range gears {
		for j := range gears {
			gears[i][j] = make([]int, 0)
		}
	}

	gear_locations := make([][]int, 0)

	valid_numbers := make([]int, 0)

	for row, line := range engine_lines {

		for col := 0; col < len(line); col++ {

			// We have found the start of a number
			if strings.Contains(numbers, string(line[col])) {

				// End of number
				end_index := col + 1
				number := string(line[col])

				for {

					// Bounds check
					if end_index > len(line)-1 {
						end_index -= 1
						break
					}

					if strings.Contains(numbers, string(line[end_index])) {

						number += string(line[end_index])
						end_index++

					} else {
						break
					}

				}

				// Now check if number is engine part

				sides := make([]side, 0)

				// Do oob checks

				if col > 0 { //Left
					sides = append(sides, Left)
				}

				if col+len(number) < len(line) { //Right
					sides = append(sides, Right)
				}

				if row > 0 { //Top
					sides = append(sides, Top)
				}

				if row < len(engine_lines)-1 { //Bottom
					sides = append(sides, Bottom)
				}

				// Now check all sides

				var valid_number bool = false
				var symbol byte = '.'
				var gear_location = []int{}

				if slices.Contains(sides, Left) && !valid_number { //Left

					if line[col-1] != '.' {
						symbol = line[col-1]
						gear_location = []int{col - 1, row}
						valid_number = true
					}

				}

				if slices.Contains(sides, Right) && !valid_number { //Right

					if line[col+len(number)] != '.' {
						symbol = line[col+len(number)]
						gear_location = []int{col + len(number), row}
						valid_number = true
					}

				}

				if slices.Contains(sides, Top) && !valid_number { //Top

					for i := col; i < col+len(number); i++ {
						if engine_lines[row-1][i] != '.' {
							symbol = engine_lines[row-1][i]
							gear_location = []int{i, row - 1}
							valid_number = true
							break
						}
					}
				}

				if slices.Contains(sides, Bottom) && !valid_number { //Top

					for i := col; i < col+len(number); i++ {
						if engine_lines[row+1][i] != '.' {
							symbol = engine_lines[row+1][i]
							gear_location = []int{i, row + 1}
							valid_number = true
							break
						}
					}
				}

				// Diagonals

				if slices.Contains(sides, Top) && (slices.Contains(sides, Left) || slices.Contains(sides, Right)) && !valid_number {

					if slices.Contains(sides, Left) {

						if engine_lines[row-1][col-1] != '.' {
							symbol = engine_lines[row-1][col-1]
							gear_location = []int{col - 1, row - 1}
							valid_number = true
						}

					}

					if slices.Contains(sides, Right) {

						if engine_lines[row-1][col+len(number)] != '.' {
							symbol = engine_lines[row-1][col+len(number)]
							gear_location = []int{col + len(number), row - 1}
							valid_number = true
						}

					}

				}

				if slices.Contains(sides, Bottom) && (slices.Contains(sides, Left) || slices.Contains(sides, Right)) && !valid_number {

					if slices.Contains(sides, Left) {

						if engine_lines[row+1][col-1] != '.' {
							symbol = engine_lines[row+1][col-1]
							gear_location = []int{col - 1, row + 1}
							valid_number = true
						}

					}

					if slices.Contains(sides, Right) {
						if engine_lines[row+1][col+len(number)] != '.' {
							symbol = engine_lines[row+1][col+len(number)]
							gear_location = []int{col + len(number), row + 1}
							valid_number = true
						}

					}
				}

				col += len(number)

				if valid_number {
					i, _ := strconv.Atoi(number)

					if symbol == '*' {

						gear_exists := false

						for _, gear := range gear_locations {
							if gear[0] == gear_location[0] && gear[1] == gear_location[1] {
								gear_exists = true
								break
							}
						}

						if !gear_exists {
							gear_locations = append(gear_locations, gear_location)
						}

						gears[gear_location[1]][gear_location[0]] = append(gears[gear_location[1]][gear_location[0]], i)

					}

					valid_numbers = append(valid_numbers, i)
				}

			}

		}

	}

	sum := 0
	ratio_sum := 0

	// P2
	gear_ratios := make([]int, 0)

	for _, gear := range gear_locations {

		if len(gears[gear[1]][gear[0]]) != 2 {
			continue
		} else {

			// I'm not proud of this line
			gear_ratios = append(gear_ratios, gears[gear[1]][gear[0]][0]*gears[gear[1]][gear[0]][1])

		}

	}

	for _, ratio := range gear_ratios {
		ratio_sum += ratio
	}

	// P1
	for _, x := range valid_numbers {
		sum += x
	}

	fmt.Println(sum)
	fmt.Println(ratio_sum)

	fmt.Printf("Time taken: %s\n", time.Since(start))

}
