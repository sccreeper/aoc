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
	Cards      string
	Score      HandType
	Bid        int
	CardCounts map[rune]int
}

//249666369

const alphabet string = "AKQJT98765432"
const alphabet_part_2 string = "AKQT98765432J"

var letter_scores map[string]int
var letter_scores_part_2 map[string]int

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

func score_hand(cards string) HandType {

	if is_five_of_a_kind(cards) {
		return FiveOfAKind
	} else if is_four_of_a_kind(cards) {
		return FourOfAKind
	} else if is_full_house(cards) {
		return FullHouse
	} else if is_three_of_a_kind(cards) {
		return ThreeOfAKind
	} else if is_two_pair(cards) {
		return TwoPair
	} else if is_one_pair(cards) {
		return OnePair
	} else {
		return High
	}

}

func parse_file(data []byte) []Hand {

	hands := make([]Hand, 0)

	lines := strings.Split(string(data), "\n")

	for _, v := range lines {

		bid, _ := strconv.Atoi(strings.Split(v, " ")[1])

		cards := strings.Split(v, " ")[0]

		card_counts := make(map[rune]int)

		for _, char := range cards {

			if _, ok := card_counts[char]; !ok {
				card_counts[char] = 1
			} else {
				card_counts[char]++
			}

		}

		hands = append(hands, Hand{
			Cards:      cards,
			Score:      score_hand(cards),
			Bid:        bid,
			CardCounts: card_counts,
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

func get_highest(cards string) string {

	highest_letter := string(cards[0])

	for _, v := range cards {
		if letter_scores_part_2[string(v)] > letter_scores_part_2[highest_letter] {
			highest_letter = string(v)
		}
	}

	return string(highest_letter)

}

func part_2(hands []Hand) int {
	// Morph hands

	morphed_hands := make([]Hand, 0)

	for _, v := range hands {

		card_frequencies := make(map[rune]int)
		for _, card := range v.Cards {
			card_frequencies[card]++
		}

		max_card := '\x00'
		max_number := 0

		for card, number := range card_frequencies {
			if number > max_number && string(card) != "J" {
				max_number = number
				max_card = card
			}
		}

		if max_number == 0 {
			morphed_hands = append(morphed_hands, v)
		} else {
			morphed_hands = append(
				morphed_hands,
				Hand{
					Cards:      strings.Replace(v.Cards, "J", string(max_card), -1),
					Score:      score_hand(strings.Replace(v.Cards, "J", string(max_card), -1)),
					Bid:        v.Bid,
					CardCounts: card_frequencies,
				},
			)
		}

	}

	hands = morphed_hands

	slices.SortStableFunc(hands, func(a, b Hand) int {

		var result int

		if a.Score == b.Score {
			for i := 0; i < len(a.Cards); i++ {
				if letter_scores_part_2[string(a.Cards[i])] > letter_scores_part_2[string(b.Cards[i])] {
					result = 1
					break
				} else if letter_scores_part_2[string(a.Cards[i])] < letter_scores_part_2[string(b.Cards[i])] {
					result = -1
					break
				}
			}
		} else if a.Score > b.Score {
			result = 1
		} else if a.Score < b.Score {
			result = -1
		}

		return result

	})

	var sum int

	for i, v := range hands {

		sum += (i + 1) * v.Bid

	}

	return sum

}

func init() {

	alphabet_reversed := reverse(alphabet)
	alphabet_part_2_reversed := reverse(alphabet_part_2)

	letter_scores = make(map[string]int)
	letter_scores_part_2 = make(map[string]int)

	for i, v := range alphabet_reversed {
		letter_scores[string(v)] = i + 1
	}

	for i, v := range alphabet_part_2_reversed {
		letter_scores_part_2[string(v)] = i + 1
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
	fmt.Println(part_2(parse_file(file_bytes)))

}
