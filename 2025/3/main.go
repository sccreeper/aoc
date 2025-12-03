package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func parse(data []byte) (result [][]int) {

	lines := strings.Split(string(data), "\n")

	result = make([][]int, 0, len(lines))

	for _, line := range lines {

		result = append(result, make([]int, 0, len(line)))

		for _, c := range line {

			x, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}

			result[len(result)-1] = append(result[len(result)-1], x)

		}

	}

	return

}

func part_one(data [][]int) (result int) {

	for _, bank := range data {

		firstNumber := 0
		firstNumberIndex := 0
		secondNumber := 0

		for i := 0; i < len(bank)-1; i++ {
			if bank[i] > firstNumber {
				firstNumber = bank[i]
				firstNumberIndex = i
			}
		}

		for i := firstNumberIndex + 1; i < len(bank); i++ {
			if bank[i] > secondNumber {
				secondNumber = bank[i]
			}
		}

		x, err := strconv.Atoi(fmt.Sprintf("%d%d", firstNumber, secondNumber))
		if err != nil {
			panic(err)
		}

		result += x

	}

	return

}

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	voltages := parse(data)

	fmt.Println(part_one(voltages))
}
