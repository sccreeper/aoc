package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Race struct {
	Time     int
	Distance int
}

func parse_part_1(data []byte) map[int]int {

	var time_row []int
	var distance_row []int

	t := strings.Split(strings.Split((strings.Split(string(data), "\n")[0]), ":")[1], " ")
	d := strings.Split(strings.Split((strings.Split(string(data), "\n")[1]), ":")[1], " ")

	for _, v := range t {
		if v == "" {
			continue
		} else {
			i, _ := strconv.Atoi(v)
			time_row = append(time_row, i)
		}
	}

	for _, v := range d {
		if v == "" {
			continue
		} else {
			i, _ := strconv.Atoi(v)
			distance_row = append(distance_row, i)
		}
	}

	result := make(map[int]int, 0)

	for i, v := range time_row {
		result[v] = distance_row[i]
	}

	return result

}

func parse_part_2(data []byte) [2]int {

	var time_number string
	var distance_number string

	t := strings.Split(strings.Split((strings.Split(string(data), "\n")[0]), ":")[1], " ")
	d := strings.Split(strings.Split((strings.Split(string(data), "\n")[1]), ":")[1], " ")

	for _, v := range t {
		if v == "" {
			continue
		} else {
			time_number += v
		}
	}

	for _, v := range d {
		if v == "" {
			continue
		} else {
			distance_number += v
		}
	}

	i, _ := strconv.Atoi(time_number)
	j, _ := strconv.Atoi(distance_number)

	return [2]int{i, j}

}

func part_1(races map[int]int) int {

	records := make(map[int]int, 0)

	for k, v := range races {

		records[k] = 0

		for i := 0; i <= k; i++ {

			if i*(k-i) > v {
				records[k]++
			}

		}

	}

	var number_of_ways int = 1

	for _, v := range records {
		number_of_ways *= v
	}

	return number_of_ways

}

func find_root(a int, b int, c int, sign int) int {

	return int(
		math.Ceil(
			(float64(-b) - (float64(sign) * (math.Sqrt(float64((b * b) - 4*a*c))))) / float64(2*a),
		),
	)

}

func part_2(race [2]int) int {

	// x . (t - x) > d
	// x . (t - x) = d
	// x . (t - x) - d = 0
	// -x^2 + tx - d = 0
	// a = -1, b = t, c = -d

	return (find_root(-1, race[0], -race[1], 1) - find_root(-1, race[0], -race[1], -1))

}

func main() {

	var file_name string
	fmt.Println("File name:")
	fmt.Scanln(&file_name)

	file_bytes, err := os.ReadFile(file_name)
	if err != nil {
		panic(err)
	}

	start := time.Now()
	fmt.Println(part_1(parse_part_1(file_bytes)))
	fmt.Println(time.Since(start))

	start = time.Now()
	fmt.Println(part_2(parse_part_2(file_bytes)))
	fmt.Println(time.Since(start))

}
