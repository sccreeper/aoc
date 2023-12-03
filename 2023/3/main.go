package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const numbers string = "1234567890"
const symbols string = "+-*/@&$#=%"

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

	engine_lines := strings.Split(string(engine_bytes), "\n")

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

				var valid_count int = 0

				if slices.Contains(sides, Left) { //Left

					if line[col-1] != '.' {
						valid_count++
					}

				}

				if slices.Contains(sides, Right) { //Right

					if line[col+len(number)] != '.' {
						valid_count++
					}

				}

				if slices.Contains(sides, Top) { //Top

					for i := col; i < col+len(number); i++ {
						if engine_lines[row-1][i] != '.' {
							valid_count++
							break
						}
					}
				}

				if slices.Contains(sides, Bottom) { //Top

					for i := col; i < col+len(number); i++ {
						if engine_lines[row+1][i] != '.' {
							valid_count++
							break
						}
					}
				}

				// Diagonals

				if slices.Contains(sides, Top) && (slices.Contains(sides, Left) || slices.Contains(sides, Right)) {

					if slices.Contains(sides, Left) {

						if engine_lines[row-1][col-1] != '.' {
							valid_count++
						}

					}

					if slices.Contains(sides, Right) {

						if engine_lines[row-1][col+len(number)] != '.' {
							valid_count++
						}

					}

				}

				if slices.Contains(sides, Bottom) && (slices.Contains(sides, Left) || slices.Contains(sides, Right)) {

					if slices.Contains(sides, Left) {

						if engine_lines[row+1][col-1] != '.' {
							valid_count++
						}

					}

					if slices.Contains(sides, Right) {
						if engine_lines[row+1][col+len(number)] != '.' {
							valid_count++
						}

					}
				}

				col += len(number)

				if valid_count > 0 {
					i, _ := strconv.Atoi(number)
					valid_numbers = append(valid_numbers, i*valid_count)
				}

			}

		}

	}

	sum := 0

	for _, x := range valid_numbers {
		sum += x
	}

	fmt.Println(sum)

}
