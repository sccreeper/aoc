package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func parse(data []byte) (result [][2]string) {

	result = make([][2]string, 0)

	for _, v := range strings.Split(string(data), ",") {

		a, b := strings.Split(v, "-")[0], strings.Split(v, "-")[1]
		result = append(result, [2]string{
			a, b,
		})

	}

	return

}

func is_invalid(val int, rep_len int) bool {
	str := strconv.Itoa(val)

	if len(str) == 2 && str[0] == str[1] {
		return true
	}

	for i := 1; i <= len(str)/rep_len; i++ {
		if len(str)%i != 0 {
			continue
		}

		substr := str[:i]
		count := strings.Count(str, substr)

		if count == rep_len && i*rep_len == len(str) {

			if str == strings.Repeat(substr, count) {
				return true
			}

		}

	}

	return false

}

func part_one(data [][2]string) int {

	var result int

	for _, v := range data {

		a, _ := strconv.Atoi(v[0])
		b, _ := strconv.Atoi(v[1])

		for i := a; i <= b; i++ {

			if is_invalid(i, 2) {
				result += i
			}

		}

	}

	fmt.Println("---")

	return result

}

func part_two(data [][2]string) int {
	var result int

	for _, v := range data {

		a, _ := strconv.Atoi(v[0])
		b, _ := strconv.Atoi(v[1])

		for i := a; i <= b; i++ {

			x := strconv.Itoa(i)

			for j := 1; j < len(x); j++ {

				if len(x)%j == 0 && strings.Repeat(x[:j], len(x)/j) == x {
					result += i
					break
				}
			}

		}

	}

	fmt.Println("---")

	return result
}

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	ids := parse(data)

	fmt.Println(part_one(ids))
	fmt.Println(part_two(ids))
}
