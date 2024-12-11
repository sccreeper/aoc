package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
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

func countStones(stones map[int]int) (total int) {
	for _, v := range stones {
		total += v
	}

	return
}

func partTwo(stones []int, blinks int) (totalStones int) {

	var stoneMap map[int]int = make(map[int]int)
	var newStoneMap map[int]int = make(map[int]int)

	// Populate dict initially

	for _, v := range stones {
		if _, exists := stoneMap[v]; exists {
			stoneMap[v]++
		} else {
			stoneMap[v] = 1
		}
	}

	for i := 0; i < blinks; i++ {

		// Order keys

		stoneKeys := make([]int, len(stoneMap))
		var j int = 0
		for k := range stoneMap {
			stoneKeys[j] = k
			j++
		}
		slices.Sort(stoneKeys)

		newStoneMap = make(map[int]int)

		for _, stone := range stoneKeys {
			total := stoneMap[stone]

			if total == 0 {
				continue
			}

			if stone == 0 {
				newStoneMap[1] += total
			} else if len(strconv.Itoa(stone))%2 == 0 {

				digits := strconv.Itoa(stone)
				numLeft, _ := strconv.Atoi(digits[:len(digits)/2])
				numRight, _ := strconv.Atoi(digits[(len(digits) / 2):])

				newStoneMap[numRight] += total
				newStoneMap[numLeft] += total

			} else {
				newK := stone * 2024
				newStoneMap[newK] += total

			}

		}

		stoneMap = newStoneMap

	}

	for _, v := range stoneMap {
		totalStones += v
	}

	return
}

func main() {

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	parsed := parseData(data)
	startTime := time.Now().UnixMilli()
	fmt.Printf("\n25 blinks P1: %d\n", partOne(parsed, 25))
	fmt.Println(time.Now().UnixMilli() - startTime)

	parsed = parseData(data)
	startTime = time.Now().UnixMilli()
	fmt.Printf("\n75 blinks P2: %d\n", partTwo(parsed, 75))
	fmt.Println(time.Now().UnixMilli() - startTime)

}
