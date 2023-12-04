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

	var cards []Card = make([]Card, len(strings.Split(cards_string, "\n")))

	for i, game := range strings.Split(cards_string, "\n") {

		card := strings.Split(game, ":")[1]
		winning := strings.Split(card, "|")[0]
		ours := strings.Split(card, "|")[1]

		cards[i] = Card{
			WinningNumbers: get_numbers(winning),
			OurNumbers:     get_numbers(ours),
		}

	}

	fmt.Println(part_2(cards))
	//fmt.Println(amounts)

}

func part_1(cards []Card) int {

	//Loop through games and do sum of powers

	var sum int

	for _, card := range cards {

		total := get_winning_amount(card)

		if total != 0 {
			sum += int(math.Pow(2, float64(total-1)))
		}

	}

	return sum

}

func get_winning_amount(card Card) int {

	total := 0

	for _, v := range card.OurNumbers {

		if slices.Contains(card.WinningNumbers, v) {
			total++
		}

	}

	return total

}

func generate_slice(start int, end int) []int {

	var slice = make([]int, 0, end-start+1)

	for i := start + 1; i <= end; i++ {
		slice = append(slice, i)
	}

	return slice

}

func part_2(cards []Card) int {

	var card_wins = make(map[int]int, 0)
	var queue []int

	for i, v := range cards {
		card_wins[i] = get_winning_amount(v)
		queue = append(queue, i)
	}

	var total int
	var queue_len int = len(queue)

	for total < queue_len {

		gen_slice := generate_slice(queue[total], queue[total]+card_wins[queue[total]])
		queue = append(queue, gen_slice...)
		total++
		queue_len += len(gen_slice)

	}

	return total

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
