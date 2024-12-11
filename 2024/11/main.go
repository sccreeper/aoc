package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
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

func partTwo(stones []int, blinks int) (totalStones int) {

	var totalChannel chan int = make(chan int)
	defer close(totalChannel)

	var wg sync.WaitGroup

	for _, v := range stones {

		go func(total chan int, val int) {

			totalChannel <- partOne([]int{val}, blinks)

		}(totalChannel, v)

	}

	wg.Wait()

	for i := 0; i < len(stones); i++ {
		totalStones += <-totalChannel
	}

	return
}

func main() {

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	parsed := parseData(data)
	fmt.Printf("Part one: %d\n", partOne(parsed, 25))
	parsed = parseData(data)
	fmt.Printf("Part two: %d\n", partTwo(parsed, 75))

}
