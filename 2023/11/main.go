package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

type Galaxy struct {
	PosX int
	PosY int
}

func gen_string(char rune, len int) (result string) {

	for i := 0; i < len; i++ {
		result += string(char)
	}

	return

}

func parse_file(data []byte, expansion_amount int) (galaxies []Galaxy, rows []int, columns []int) {

	galaxies = make([]Galaxy, 0)

	lines := strings.Split(string(data), "\n")

	//Find rows & columns with no galaxies

	rows = make([]int, 0)
	columns = make([]int, 0)

	for y := 0; y < len(lines); y++ {

		galaxy_found := false

		for x := 0; x < len(lines[0]); x++ {

			if lines[y][x] == '#' {
				galaxy_found = true
				break
			}

		}

		if !galaxy_found {
			rows = append(rows, y)
		}

	}

	for x := 0; x < len(lines[0]); x++ {

		galaxy_found := false

		for y := 0; y < len(lines); y++ {

			if lines[y][x] == '#' {
				galaxy_found = true
				break
			}

		}

		if !galaxy_found {
			columns = append(columns, x)
		}

	}

	//Insert additional rows and columns

	row_lines := make([]string, 0)

	for i, v := range lines {

		if !slices.Contains(rows, i) {
			row_lines = append(row_lines, v)
		} else {
			row_lines = append(row_lines, v)

			for j := 0; j < expansion_amount; j++ {
				row_lines = append(row_lines, gen_string('.', len(lines[0])))
			}
		}

	}

	column_lines := make([]string, 0)

	for y := 0; y < len(row_lines); y++ {

		line := ""

		for x := 0; x < len(row_lines[0]); x++ {

			if !slices.Contains(columns, x) {
				line += string(row_lines[y][x])
			} else {
				line += string(row_lines[y][x])

				for j := 0; j < expansion_amount; j++ {
					line += "."
				}
			}

		}

		column_lines = append(column_lines, line)

	}

	//Find galaxies

	for y := 0; y < len(column_lines); y++ {

		for x := 0; x < len(column_lines[0]); x++ {

			if column_lines[y][x] == '#' {
				galaxies = append(galaxies, Galaxy{PosX: x, PosY: y})
			}

		}

	}

	return

}

func get_distance(a Galaxy, b Galaxy) int {

	return int(math.Abs(float64(a.PosX-b.PosX)) + math.Abs(float64(a.PosY-b.PosY)))

}

func part_1(galaxies []Galaxy) (sum int) {

	sum = 0

	done_galaxies := make([][]int, len(galaxies))

	for i := 0; i < len(galaxies); i++ {
		done_galaxies[i] = make([]int, 0)
	}

	for j, v := range galaxies {

		for i := 0; i < len(galaxies); i++ {

			if galaxies[i].PosX == v.PosX && galaxies[i].PosY == v.PosY {
				continue
			} else {

				if !slices.Contains(done_galaxies[i], j) {
					done_galaxies[j] = append(done_galaxies[j], i)

					sum += get_distance(v, galaxies[i])
				}

			}

		}

	}

	return

}

func part_2(galaxies []Galaxy, rows []int, columns []int) (sum int) {

	var transform int = int(math.Pow10(6)) - 1

	sum = 0

	done_galaxies := make([][]int, len(galaxies))

	for i := 0; i < len(galaxies); i++ {
		done_galaxies[i] = make([]int, 0)
	}

	for j, v := range galaxies {

		for i := 0; i < len(galaxies); i++ {

			if galaxies[i].PosX == v.PosX && galaxies[i].PosY == v.PosY {
				continue
			} else {

				if !slices.Contains(done_galaxies[i], j) {
					done_galaxies[j] = append(done_galaxies[j], i)

					//Calculate distance

					scale_factor := 0

					for _, row := range rows {
						if v.PosY < row && galaxies[i].PosY > row {
							scale_factor++
						} else if v.PosY > galaxies[i].PosY && v.PosY > row && galaxies[i].PosY < row {
							scale_factor++
						}
					}

					for _, column := range columns {
						if v.PosX < column && galaxies[i].PosX > column {
							scale_factor++
						} else if v.PosX > galaxies[i].PosX && v.PosX > column && galaxies[i].PosX < column {
							scale_factor++
						}
					}

					sum += get_distance(v, galaxies[i]) + (transform * scale_factor)
				}

			}

		}

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

	galaxies, _, _ := parse_file(file_data, 1)
	fmt.Println(part_1(galaxies))

	fmt.Println(part_2(parse_file(file_data, 0)))
	//fmt.Println(part_1(parse_file(file_data, int(math.Pow10(6)-1))))

}
