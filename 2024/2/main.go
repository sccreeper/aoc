package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

var saveLevels []int = []int{
	-1, -2, -3, 1, 2, 3,
}

type changeDirection int

const increasing changeDirection = 0
const decreasing changeDirection = 1

func parseLine(line string) (result []int) {

	result = make([]int, 0)

	strNums := strings.Split(line, " ")

	for _, v := range strNums {

		num, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		result = append(result, num)

	}

	return

}

func parseData(data []byte) (result [][]int) {

	accumulator := ""

	result = make([][]int, 0)

	for _, b := range data {

		if b == '\n' {

			result = append(result, parseLine(accumulator))
			accumulator = ""

		} else {
			accumulator += string(b)
		}

	}

	result = append(result, parseLine(accumulator))

	return

}

func safeCount(line []int) int {

	var direction changeDirection
	var safeCount int = 0

	if line[0] > line[1] {
		direction = decreasing
	} else if line[0] < line[1] {
		direction = increasing
	} else if line[0] == line[1] {
		safeCount++
	}

	for j := 1; j < len(line); j++ {

		if line[j] > line[j-1] {

			if direction != increasing {
				safeCount++
			}

		} else if line[j] < line[j-1] {

			if direction != decreasing {
				safeCount++
			}

		}

		if !slices.Contains(saveLevels, line[j]-line[j-1]) {
			safeCount++
		}

	}

	return safeCount
}

func partOne(data []byte) {

	parsedData := parseData(data)

	safeTotal := 0

	for i := 0; i < len(parsedData); i++ {

		if safeCount(parsedData[i]) == 0 {
			safeTotal++
		}
	}

	fmt.Println(safeTotal)

}

func partTwo(data []byte) {

	parsed := parseData(data)

	var safeTotal int = 0

	for i := 0; i < len(parsed); i++ {

		if safeCount(parsed[i]) == 0 {
			// fmt.Println(parsed[i])
			safeTotal++
		} else {

			for j := 0; j < len(parsed[i]); j++ {

				new := make([]int, 0)
				new = append(new, parsed[i][:j]...)
				new = append(new, parsed[i][j+1:]...)

				if safeCount(new) == 0 {
					// fmt.Println(parsed[i])
					safeTotal++
					break
				}

			}

		}

	}

	fmt.Println(safeTotal)

}

func main() {
	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	partOne(bytes)
	partTwo(bytes)

}
