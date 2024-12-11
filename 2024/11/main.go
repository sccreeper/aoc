package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func parseData(data []byte) (stones []int) {

	var stonesString []string = strings.Split(string(data), " ")
	stones = make([]int, 0)

	for _, v := range stonesString {
		n, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		stones = append(stones, n)
	}

	return

}

func partOne(stones []int, blinks int) (totalStones int) {

	for i := 0; i < blinks; i++ {

		fmt.Printf("\rBlink %d/%d...", i+1, blinks)

		for j := 0; j < len(stones); j++ {

			if stones[j] == 0 {
				stones[j] = 1
			} else if len(strconv.Itoa(stones[j]))%2 == 0 {

				digits := strconv.Itoa(stones[j])
				numLeft, _ := strconv.Atoi(digits[:len(digits)/2])
				numRight, _ := strconv.Atoi(digits[(len(digits) / 2):])

				stones[j] = numLeft

				if j+1 >= len(stones) {
					stones = append(stones, numRight)
				} else {
					stones = slices.Insert(stones, j+1, numRight)
				}

				j++

			} else {
				stones[j] *= 2024
			}

		}

	}

	totalStones = len(stones)

	return
}

func main() {

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	parsed := parseData(data)
	fmt.Printf("Part one: %d\n", partOne(parsed, 25))

}
