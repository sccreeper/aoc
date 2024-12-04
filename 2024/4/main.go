package main

import (
	"fmt"
	"io"
	"os"
	"slices"
)

var xmasString = "XMAS"
var masString = "MAS"
var xmasStringRev string

func init() {
	xmasStringRune := []rune(xmasString)
	slices.Reverse(xmasStringRune)
	xmasStringRev = string(xmasStringRune)

}

type Pos struct {
	X int
	Y int
}

func parseData(data []byte) (result [][]rune) {

	var accumulator string

	result = make([][]rune, 0)

	for _, v := range data {

		if v == '\n' {
			result = append(result, []rune(accumulator))
			accumulator = ""
		} else {
			accumulator += string(v)
		}

	}

	result = append(result, []rune(accumulator))

	return

}

func search2d(p Pos, data [][]rune, searchVector Pos, searchString string) (found int) {

	var valid bool = true
	var charIndex = 0
	var foundString string = ""

	for charIndex < len(searchString) && valid && p.X < len(data[0]) && p.Y < len(data) && p.X >= 0 && p.Y >= 0 {

		if data[p.Y][p.X] != rune(searchString[charIndex]) {
			valid = false
			break
		}

		foundString += string(data[p.Y][p.X])
		charIndex++

		p.X += searchVector.X
		p.Y += searchVector.Y

	}

	if foundString == searchString {

		fmt.Printf("X: %d Y: %d\n", p.X-(searchVector.X*len(searchString)), p.Y-(searchVector.Y*len(searchString)))

		return 1
	} else {
		return 0
	}

}

// func searchHorizontally(p Pos, data [][]rune) (matches int) {
// 	// ->
// 	if !(p.X+len(xmasString) > len(data[0])) && string(data[p.Y][p.X:p.X+len(xmasString)]) == xmasString {
// 		matches++
// 	}

// 	// <-
// 	if !(p.X-len(xmasString)+1 < 0) && string(data[p.Y][p.X-len(xmasString)+1:p.X+1]) == xmasStringRev {
// 		matches++
// 	}

// 	return

// }

func partOne(data [][]rune) (matches int) {

	matches = 0

	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[0]); x++ {

			if data[y][x] == 'X' {

				// Horizontal
				matches += search2d(
					Pos{X: x, Y: y},
					data,
					Pos{X: -1, Y: 0},
					xmasString,
				)

				matches += search2d(
					Pos{X: x, Y: y},
					data,
					Pos{X: 1, Y: 0},
					xmasString,
				)

				// Vertical
				matches += search2d(
					Pos{X: x, Y: y},
					data,
					Pos{X: 0, Y: 1},
					xmasString,
				)

				matches += search2d(
					Pos{X: x, Y: y},
					data,
					Pos{X: 0, Y: -1},
					xmasString,
				)

				// Diagonal

				matches += search2d(
					Pos{X: x, Y: y},
					data,
					Pos{X: -1, Y: -1},
					xmasString,
				)

				matches += search2d(
					Pos{X: x, Y: y},
					data,
					Pos{X: 1, Y: 1},
					xmasString,
				)

				matches += search2d(
					Pos{X: x, Y: y},
					data,
					Pos{X: -1, Y: 1},
					xmasString,
				)

				matches += search2d(
					Pos{X: x, Y: y},
					data,
					Pos{X: 1, Y: -1},
					xmasString,
				)

			}

		}
	}

	return

}

// Couldn't get this to work with my search2d code so I just did this instead.
func partTwo(data [][]rune) (matches int) {
	matches = 0

	for y := 1; y < len(data)-1; y++ {
		for x := 1; x < len(data[0])-1; x++ {

			if data[y][x] != 'A' {
				continue
			}

			topLeftChar := data[y-1][x-1]
			topRightChar := data[y-1][x+1]
			bottomLeftChar := data[y+1][x-1]
			bottomRightChar := data[y+1][x+1]

			// Compare strings

			compA := string([]rune{topLeftChar, 'A', bottomRightChar})
			compB := string([]rune{topRightChar, 'A', bottomLeftChar})

			if compA != "MAS" && compA != "SAM" {
				continue
			}

			if compB != "MAS" && compB != "SAM" {
				continue
			}

			matches++

		}
	}

	return
}

func main() {
	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	parsed := parseData(bytes)
	fmt.Println(partOne(parsed))
	fmt.Println(partTwo(parsed))
}
