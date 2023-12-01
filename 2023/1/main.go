package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const integers = "1234567890"

var words []string = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var word_lengths []int = []int{3, 3, 5, 4, 4, 3, 5, 5, 4}
var word_digits []string = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

//go:embed calibration.txt
var input_string string
var total int

func main() {

	start := time.Now()

	for _, line := range strings.Split(input_string, "\n") {

		var line_numbers []string = make([]string, 0)
		var line_length int = len(line)

		for i := 0; i < line_length; i++ {

			if strings.Contains(integers, string(line[i])) { //Words can't start with integers.

				line_numbers = append(line_numbers, string(line[i]))

			} else {

				for j := 0; j < 9; j++ {

					if i+word_lengths[j] < line_length+1 {

						if string(line[i:i+word_lengths[j]]) == words[j] {

							line_numbers = append(line_numbers, word_digits[j])

						}

					}

				}

			}

		}

		if len(line_numbers) == 0 {
			continue
		} else if len(line_numbers) == 1 {

			i, _ := strconv.Atoi(line_numbers[0] + line_numbers[0])
			total += i

		} else {

			i, _ := strconv.Atoi(line_numbers[0] + line_numbers[len(line_numbers)-1])
			total += i

		}

	}

	used := time.Since(start)
	fmt.Println("Time used:", used)

	fmt.Println(total)

}
