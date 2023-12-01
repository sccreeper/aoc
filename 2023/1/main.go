package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const integers = "1234567890"

var words []string = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var word_digits []string = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

func main() {

	var input_name string
	fmt.Println("Input file name:")
	fmt.Scanln(&input_name)

	input_data, err := os.ReadFile(input_name)
	if err != nil {
		fmt.Printf("No input called: %s", input_name)
	}

	input_string := string(input_data)

	var total int

	for _, line := range strings.Split(input_string, "\n") {

		var line_numbers []string = make([]string, 0)

		for i := 0; i < len(line); i++ {

			if strings.Contains(integers, string(line[i])) { //Words can't start with integers.

				line_numbers = append(line_numbers, string(line[i]))

			} else {

				for j := 0; j < len(words); j++ {

					if i+len(words[j]) < len(line)+1 {

						if string(line[i:i+len(words[j])]) == words[j] {

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
			fmt.Println(i)

		} else {

			i, _ := strconv.Atoi(line_numbers[0] + line_numbers[len(line_numbers)-1])
			total += i
			fmt.Println(i)

		}

	}

	fmt.Println(total)

}
