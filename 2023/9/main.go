package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func parse_file(data []byte) (sequences [][]int) {

	sequences = make([][]int, 0)

	for i, v := range strings.Split(string(data), "\n") {

		sequences = append(sequences, make([]int, 0))

		numbers := strings.Split(v, " ")

		for _, n := range numbers {

			x, _ := strconv.Atoi(n)

			sequences[i] = append(sequences[i], x)

		}

	}

	return

}

func get_differences(numbers []int) (result []int) {

	result = make([]int, 0)

	for i := 1; i < len(numbers); i++ {

		result = append(result, numbers[i]-numbers[i-1])

	}

	return

}

func all_same(numbers []int) bool {

	x := numbers[0]

	for _, v := range numbers {
		if v != x {
			return false
		}
	}

	return true

}

func get_sequence(numbers []int) (all_differences [][]int) {

	// Determine degree

	var differences_array []int = numbers

	all_differences = make([][]int, 0)

	for {

		if all_same(differences_array) {
			break
		}

		differences_array = get_differences(differences_array)
		all_differences = append(all_differences, differences_array)

	}

	return
}

func part_1(sequences [][]int) int {

	var sum int = 0

	for _, s := range sequences {

		// Get amount to add

		diffs := get_sequence(s)
		slices.Reverse(diffs)

		total := 0

		for i := 0; i < len(diffs); i++ {
			total += diffs[i][len(diffs[i])-1]
		}

		sum += s[len(s)-1] + total

	}

	return sum

}

func part_2(sequences [][]int) int {

	var sum int = 0

	for _, s := range sequences {

		// Get amount to add

		diffs := get_sequence(s)
		slices.Reverse(diffs)

		total := 0

		for i := 0; i < len(diffs); i++ {
			total = diffs[i][0] - total
		}

		sum += s[0] - total

	}

	return sum

}

func main() {

	var file_name string
	fmt.Println("File name:")
	fmt.Scanln(&file_name)

	file_data, err := os.ReadFile(file_name)
	if err != nil {
		panic(err)
	}

	start := time.Now()

	fmt.Println(part_1(parse_file(file_data)))
	fmt.Println(part_2(parse_file(file_data)))

	fmt.Println(time.Since(start))

}
