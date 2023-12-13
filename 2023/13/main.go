package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

type Block struct {
	Text    string
	Lines   []string
	Columns []string
}

func parse_file(data []byte) (result []Block) {

	result = make([]Block, 0)
	var current_block string

	for _, v := range strings.Split(string(data), "\n") {

		if v == "" || v == "\n" {

			result = append(result, Block{Text: current_block, Lines: strings.Split(current_block, "\n")})
			current_block = ""

		} else {
			current_block += v + "\n"
		}

	}

	result = append(result, Block{Text: current_block, Lines: strings.Split(current_block, "\n")})

	//Process columns

	for i, block := range result {

		for j := 0; j < len(block.Lines[0]); j++ {

			var current_column string

			for k := 0; k < len(block.Lines); k++ {
				if block.Lines[k] == "" {
					continue
				}

				current_column += string(block.Lines[k][j])
			}

			result[i].Columns = append(result[i].Columns, current_column)

		}

	}

	fmt.Printf("Found %d blocks\n", len(result))

	return

}

func check_reflection(data []string, index int) bool {

	var scan_distance int

	if index >= int(len(data)/2) {
		scan_distance = len(data) - (index + 1)
	} else if index != len(data)/2 {
		scan_distance = (index + 1)
	} else {
		scan_distance = index
	}

	if len(data) == 0 {
		return false
	}

	if index-scan_distance < 0 {
		return false
	}

	left := data[index-scan_distance : index]
	right := data[index : index+scan_distance]
	slices.Reverse(right)

	if len(right) == 0 || len(left) == 0 {
		return false
	}

	for i := 0; i < len(left); i++ {

		if left[i] != right[i] {
			return false
		}

	}

	return true

}

func find_reflection(data []string) (index int, found bool) {

	var previous string = data[0]

	for i := 1; i < len(data); i++ {

		if data[i] == previous && check_reflection(data, i) {
			index = i
			found = true
			return
		}

		previous = data[i]

	}

	found = false

	return

}

func part_1(blocks []Block) (sum int) {

	sum = 0

	for _, block := range blocks {

		//Loop through lines then columns to try and find previous matches

		if index, exists := find_reflection(block.Lines); exists {
			sum += 100 * index
			continue
			//fmt.Println("l")
			//fmt.Println(i)
		}

		if index, exists := find_reflection(block.Columns); exists {
			//fmt.Println("j")
			//fmt.Println(i)
			sum += index
		}

	}

	return sum

}

func main() {

	var file_name string
	fmt.Println("Enter file name:")
	fmt.Scanln(&file_name)

	file_data, err := os.ReadFile(file_name)
	if err != nil {
		panic(err)
	}

	fmt.Println(part_1(parse_file(file_data)))

}
