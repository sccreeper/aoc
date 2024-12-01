package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parseLine(line string) [2]int {
	nums := strings.Split(line, " ")
	numsClean := []string{nums[0], nums[len(nums)-1]}

	numLeft, err := strconv.Atoi(numsClean[0])
	if err != nil {
		panic(err)
	}

	numRight, err := strconv.Atoi(numsClean[1])
	if err != nil {
		panic(err)
	}

	return [2]int{numLeft, numRight}
}

func parseData(data []byte) ([]int, []int) {

	arrayLeft := make([]int, 0)
	arrayRight := make([]int, 0)

	// Parse input

	accumulator := ""

	for _, b := range data {

		if b == '\n' {

			nums := parseLine(accumulator)

			arrayLeft = append(arrayLeft, nums[0])
			arrayRight = append(arrayRight, nums[1])

			accumulator = ""

		} else {
			accumulator += string(b)
		}

	}

	nums := parseLine(accumulator)

	arrayLeft = append(arrayLeft, nums[0])
	arrayRight = append(arrayRight, nums[1])

	return arrayLeft, arrayRight

}

func partOne(data []byte) {

	arrayLeft, arrayRight := parseData(data)

	slices.Sort(arrayLeft)
	slices.Sort(arrayRight)

	var total int

	for i := 0; i < len(arrayRight); i++ {

		total += int(math.Abs(float64(arrayRight[i] - arrayLeft[i])))

	}

	fmt.Println(total)

}

func partTwo(data []byte) {

	arrayLeft, arrayRight := parseData(data)

	simMap := make(map[int]int, 0)

	var total int

	for _, v := range arrayLeft {

		val, exists := simMap[v]

		if exists {
			total += val
		} else {

			var numTimes int

			for _, n := range arrayRight {
				if n == v {
					numTimes++
				}
			}

			simMap[v] = numTimes * v
			total += simMap[v]

		}

	}

	fmt.Println(total)

}

func main() {
	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	partOne(bytes)
	partTwo(bytes)
}
