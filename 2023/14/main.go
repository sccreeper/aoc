package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Tile rune

const (
	CubeRock  Tile = '#'
	RoundRock Tile = 'O'
	Empty     Tile = '.'
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

func part_1(tiles [][]Tile) (sum int) {

	var moved int = 1

	//Move rocks

	for moved != 0 {

		moved = 0

		for i := 0; i < len(tiles); i++ {

			for j := 0; j < len(tiles[0]); j++ {

				if tiles[i][j] == RoundRock && i-1 >= 0 && tiles[i-1][j] == Empty {

					tiles[i-1][j] = RoundRock
					tiles[i][j] = Empty

					moved++

				} else {
					continue
				}

			}

		}

	}

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

func main() {

	var file_name string
	fmt.Println("File name:")
	fmt.Scanln(&file_name)

	file_data, err := os.ReadFile(file_name)
	if err != nil {
		panic(err)
	}

	fmt.Println(part_1(parse_file(file_data)))

}
