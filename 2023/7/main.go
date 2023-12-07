package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type HandType int

const (
	FiveOfAKind  HandType = 7
	FourOfAKind  HandType = 6
	FullHouse    HandType = 5
	ThreeOfAKind HandType = 4
	TwoPair      HandType = 3
	OnePair      HandType = 2
	High         HandType = 1
)

type Hand struct {
	Cards string
	Score HandType
	Bid   int
}

const alphabet string = "AKQJT98765432"

var letter_scores map[string]int

func n_of_a_kind(cards string, number int) bool {

	card_counts := make(map[byte]int, 0)

	for _, v := range cards {

		if _, ok := card_counts[byte(v)]; !ok {
			card_counts[byte(v)] = 1
		} else {
			card_counts[byte(v)]++
		}

	}

	for _, v := range card_counts {
		if v == number {
			return true
		}
	}

	return false

}

func n_pair(cards string, n_pairs int) bool {

	card_counts := make(map[byte]int, 0)

	for _, v := range cards {

		if _, ok := card_counts[byte(v)]; !ok {
			card_counts[byte(v)] = 1
		} else {
			card_counts[byte(v)]++
		}

	}

	pairs := 0

	for _, v := range card_counts {
		if v == 2 {
			pairs++
		}
	}

	return pairs == n_pairs

}

func is_five_of_a_kind(cards string) bool {

	for _, v := range cards {

		if byte(v) != cards[0] {
			return false
		}

	}

	return true
}

func is_four_of_a_kind(cards string) bool {
	return n_of_a_kind(cards, 4)
}

func is_full_house(cards string) bool {

	card_counts := make(map[byte]int, 0)

	for _, v := range cards {

		if _, ok := card_counts[byte(v)]; !ok {
			card_counts[byte(v)] = 1
		} else {
			card_counts[byte(v)]++
		}

	}

	return len(card_counts) == 2

}

func is_three_of_a_kind(cards string) bool {
	return n_of_a_kind(cards, 3)
}

func is_two_pair(cards string) bool {
	return n_pair(cards, 2)
}

func is_one_pair(cards string) bool {
	return n_pair(cards, 1)
}

func parse_file(data []byte) []Hand {

	hands := make([]Hand, 0)

	lines := strings.Split(string(data), "\n")

	for _, v := range lines {

		bid, _ := strconv.Atoi(strings.Split(v, " ")[1])

		var hand_type HandType
		cards := strings.Split(v, " ")[0]

		if is_five_of_a_kind(cards) {
			hand_type = FiveOfAKind
		} else if is_four_of_a_kind(cards) {
			hand_type = FourOfAKind
		} else if is_full_house(cards) {
			hand_type = FullHouse
		} else if is_three_of_a_kind(cards) {
			hand_type = ThreeOfAKind
		} else if is_two_pair(cards) {
			hand_type = TwoPair
		} else if is_one_pair(cards) {
			hand_type = OnePair
		} else {
			hand_type = High
		}

		hands = append(hands, Hand{
			Cards: cards,
			Score: hand_type,
			Bid:   bid,
		})

	}

	return hands

}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func part_1(hands []Hand) int {

	sort.Slice(hands, func(i, j int) bool {

		if hands[i].Score == hands[j].Score {

			//Sort by character scores

			for k := 0; k < 5; k++ {

				if hands[i].Cards[k] == hands[j].Cards[k] {
					continue
				} else {
					return letter_scores[string(hands[i].Cards[k])] < letter_scores[string(hands[j].Cards[k])]
				}

			}

		}

		return hands[i].Score < hands[j].Score

	})

	var sum int

	slices.Reverse(hands)

	for i, v := range hands {

		sum += (len(hands) - (i)) * v.Bid
		//fmt.Printf("%d * %d\n", len(hands)-i, v.Bid)

	}

	return sum
}

func init() {

	alphabet_reversed := reverse(alphabet)

	letter_scores = make(map[string]int)

	for i, v := range alphabet_reversed {
		letter_scores[string(v)] = i + 1
	}

}

func main() {

	var file_name string
	fmt.Println("File name")
	fmt.Scanln(&file_name)

	file_bytes, err := os.ReadFile(file_name)
	if err != nil {
		panic(err)
	}

	fmt.Println(part_1(parse_file(file_bytes)))

}
