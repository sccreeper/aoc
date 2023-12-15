package main

import (
	"fmt"
	"os"
	"strings"
)

func parse_file(data []byte) []string {

	return strings.Split(string(data), ",")

}

func part_1(data []string) (sum int) {

	sum = 0

	for _, v := range data {

		current_sum := 0

		for _, b := range []byte(v) {

			current_sum += int(b)
			current_sum *= 17
			current_sum %= 256

		}

		sum += current_sum

	}

	return

}

func main() {

	var file_name string
	fmt.Println("File name:")
	fmt.Scanln(&file_name)

	file_data, err := os.ReadFile(file_name)
	if err != nil {
		panic(err)
	}

	fmt.Println(part_1(parse_file(file_data)))

}
