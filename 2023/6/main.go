package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
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

func part_2(race [2]int) int {

	var records int

	for i := 0; i < race[0]; i++ {

		if i*(race[0]-i) > race[1] {
			records++
		}

	}

	return records

}

func main() {

	var file_name string
	fmt.Println("File name:")
	fmt.Scanln(&file_name)

	file_bytes, err := os.ReadFile(file_name)
	if err != nil {
		panic(err)
	}

	fmt.Println(part_1(parse_part_1(file_bytes)))
	fmt.Println(part_2(parse_part_2(file_bytes)))

}
