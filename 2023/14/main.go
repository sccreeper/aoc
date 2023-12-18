package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Tile rune
type Direction byte

const (
	CubeRock  Tile = '#'
	RoundRock Tile = 'O'
	Empty     Tile = '.'
)

const (
	North Direction = 0
	South Direction = 1
	East  Direction = 2
	West  Direction = 3
)

func parse_file(data []byte) (tiles [][]Tile) {

	tiles = make([][]Tile, len(strings.Split(string(data), "\n")))

	for i, line := range strings.Split(string(data), "\n") {

		tiles[i] = make([]Tile, 0)

		for _, char := range line {

			switch char {
			case '#':
				tiles[i] = append(tiles[i], CubeRock)
			case 'O':
				tiles[i] = append(tiles[i], RoundRock)
			case '.':
				tiles[i] = append(tiles[i], Empty)
			}

		}

	}

	return

}

func oob_check(dir Direction, i int, j int, w int, h int) bool {

	switch dir {
	case North:
		return i-1 < 0
	case South:
		return i+1 >= h
	case East:
		return j-1 < 0
	case West:
		return i+1 >= w
	default:
		return false
	}

}

func empty_check(dir Direction, i int, j int) (i_new int, j_new int) {

	switch dir {
	case North:
		i_new = i - 1
		j_new = j
		return
	case South:
		i_new = i + 1
		j_new = j
		return
	case East:
		i_new = i
		j_new = j - 1
		return
	case West:
		i_new = i
		j_new = j + 1
		return
	default:
		return
	}

}

func tilt(dir Direction, tiles [][]Tile) [][]Tile {

	var moved int = 1

	//Move rocks

	for moved != 0 {

		moved = 0

		for i := 0; i < len(tiles); i++ {

			for j := 0; j < len(tiles[0]); j++ {

				new_i, new_j := empty_check(dir, i, j)

				if tiles[i][j] == RoundRock && !oob_check(dir, i, j, len(tiles[0]), len(tiles)) && tiles[new_i][new_j] == Empty {

					tiles[new_i][new_j] = RoundRock
					tiles[i][j] = Empty

					moved++

				} else {
					continue
				}

			}

		}

	}

	return tiles

}

func calculate_load(tiles [][]Tile) (sum int) {

	//Calculate load

	slices.Reverse(tiles)

	for i := 0; i < len(tiles); i++ {

		for j := 0; j < len(tiles[0]); j++ {

			if tiles[i][j] == RoundRock {

				sum += i + 1

			}

		}

	}

	return

}

func part_1(tiles [][]Tile) (sum int) {

	tiles = tilt(North, tiles)
	sum = calculate_load(tiles)

	return

}

func part_2(tiles [][]Tile) (load int) {

	for i := 0; i < 1000000000; i++ {

		tiles = tilt(North, tiles)
		tiles = tilt(West, tiles)
		tiles = tilt(East, tiles)
		tiles = tilt(South, tiles)

	}

	load = calculate_load(tiles)

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
