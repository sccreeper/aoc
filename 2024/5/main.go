package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parseData(data []byte) (orderingRules map[int][]int, updateNumbers [][]int) {

	lines := make([]string, 0)

	// Parse lines

	accumulator := ""
	dataSplitIndex := 0

	for _, b := range data {

		if b == '\n' {

			lines = append(lines, accumulator)

			if len(accumulator) == 0 {
				dataSplitIndex = len(lines) - 1
			}

			accumulator = ""

		} else {
			accumulator += string(b)
		}

	}

	lines = append(lines, accumulator)

	orderingRulesString := lines[:dataSplitIndex]
	pageUpdateNumbersString := lines[dataSplitIndex+1:]

	orderingRules = make(map[int][]int)

	for _, v := range orderingRulesString {

		nums := strings.Split(v, "|")

		numKey, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}

		numVal, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}

		if _, exists := orderingRules[numKey]; exists {

			orderingRules[numKey] = append(orderingRules[numKey], numVal)

		} else {

			orderingRules[numKey] = make([]int, 0)
			orderingRules[numKey] = append(orderingRules[numKey], numVal)

		}

	}

	updateNumbers = make([][]int, 0)

	for i, v := range pageUpdateNumbersString {

		nums := strings.Split(v, ",")
		updateNumbers = append(updateNumbers, make([]int, 0))

		for _, n := range nums {

			numConv, err := strconv.Atoi(n)
			if err != nil {
				panic(err)
			}

			updateNumbers[i] = append(updateNumbers[i], numConv)

		}

	}

	return

}

func hasIntersection(sliceA []int, sliceB []int) bool {

	for _, v := range sliceA {

		if slices.Contains(sliceB, v) {
			return true
		}

	}

	return false

}

func partOne(orderingRules map[int][]int, updateNumbers [][]int) (result int) {

	correctlyOrdered := make([][]int, 0)

	for _, v := range updateNumbers {

		valid := true

		for i := 1; i < len(v); i++ {

			if hasIntersection(v[:i], orderingRules[v[i]]) {
				valid = false
				break
			}

		}

		if valid {
			correctlyOrdered = append(correctlyOrdered, v)
		}

	}

	// Get middle numbers and add

	for _, v := range correctlyOrdered {

		result += v[len(v)/2]

	}

	return
}

func partTwo(orderingRules map[int][]int, updateNumbers [][]int) (result int) {

	incorrectlyOrdered := make([][]int, 0)

	for _, v := range updateNumbers {

		for i := 1; i < len(v); i++ {

			if hasIntersection(v[:i], orderingRules[v[i]]) {
				incorrectlyOrdered = append(incorrectlyOrdered, v)
				break
			}

		}

	}

	// Order the incorrect arrays

	for i := 0; i < len(incorrectlyOrdered); i++ {

		var hasIntersections = true

		for hasIntersections {

			var j int = 0

			for j < len(incorrectlyOrdered[i]) {

				if hasIntersection(incorrectlyOrdered[i][:j], orderingRules[incorrectlyOrdered[i][j]]) {

					// Get indexes it needs to be before

					badIndexes := make([]int, 0)

					for _, k := range orderingRules[incorrectlyOrdered[i][j]] {

						idx := slices.Index(incorrectlyOrdered[i], k)
						if idx != -1 {
							badIndexes = append(badIndexes, idx)
						}

					}

					var moveIndex int
					slices.Sort(badIndexes)
					moveIndex = badIndexes[0]

					moveNum := incorrectlyOrdered[i][j]

					// Remove item
					incorrectlyOrdered[i] = append(incorrectlyOrdered[i][:j], incorrectlyOrdered[i][j+1:]...)

					// Add item at "correct index"

					incorrectlyOrdered[i] = append(
						incorrectlyOrdered[i][:moveIndex],
						append(
							[]int{moveNum},
							incorrectlyOrdered[i][moveIndex:]...,
						)...,
					)

					j = 0

				} else {
					j++
				}

			}

			hasIntersections = false

		}

	}

	// Get middle numbers and add

	for _, v := range incorrectlyOrdered {

		result += v[len(v)/2]

	}

	return

}

func main() {

	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	orderingRules, updateNumbers := parseData(bytes)
	fmt.Println(partOne(orderingRules, updateNumbers))
	fmt.Println(partTwo(orderingRules, updateNumbers))

}
