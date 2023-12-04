package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Card struct {
	WinningNumbers []int
	OurNumbers     []int
}

func main() {

	var file_name string
	fmt.Println("Enter file name:")
	fmt.Scanln(&file_name)

	file_bytes, err := os.ReadFile(file_name)
	if err != nil {
		panic(err)
	}

	//Parse

	cards_string := string(file_bytes)

	var cards []Card = make([]Card, len(cards_string))

	for i, game := range strings.Split(cards_string, "\n") {

		card := strings.Split(game, ":")[1]
		winning := strings.Split(card, "|")[0]
		ours := strings.Split(card, "|")[1]

		cards[i] = Card{
			WinningNumbers: get_numbers(winning),
			OurNumbers:     get_numbers(ours),
		}

	}

	//Loop through games and do sum of powers

	var sum int

	for _, card := range cards {

		total := 0

		for _, v := range card.OurNumbers {

			if slices.Contains(card.WinningNumbers, v) {
				total += 1
			}

		}

		if total != 0 {
			sum += int(math.Pow(2, float64(total-1)))
		}

	}

	fmt.Println(sum)

}

func get_numbers(numbers string) []int {

	var result []int = make([]int, 0)

	for _, v := range strings.Split(numbers, " ") {

		if v == "" {
			continue
		}

		i, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		result = append(result, i)

	}

	return result

}
