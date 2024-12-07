package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type instruction struct {
	Target     int
	Components []int
}

func parseData(data []byte) (result []instruction) {

	lines := make([]string, 0)

	accumulator := ""

	for _, b := range data {

		if b == '\n' {
			lines = append(lines, accumulator)
			accumulator = ""
		} else {
			accumulator += string(b)
		}

	}

	lines = append(lines, accumulator)

	result = make([]instruction, 0, len(lines))

	for i, line := range lines {

		target := strings.Split(line, ":")[0]
		nums := strings.Split(line, ":")[1]
		nums = strings.TrimSpace(nums)

		result = append(result, instruction{})

		targetNum, err := strconv.Atoi(target)
		if err != nil {
			panic(err)
		}

		result[i].Target = targetNum
		result[i].Components = make([]int, 0)

		for _, v := range strings.Split(nums, " ") {

			n, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}

			result[i].Components = append(result[i].Components, n)

		}

	}

	return

}

func isTruePartOne(instr instruction) bool {

	var possPermut uint64 = uint64(math.Pow(2.0, float64(len(instr.Components)-1)))

	var i uint64
	for i = 0; i < possPermut; i++ {

		sum := instr.Components[0]

		for j := 1; j < len(instr.Components); j++ {

			operator := (i >> (j - 1)) & 1

			if operator == 1 {
				sum *= instr.Components[j]
			} else {
				sum += instr.Components[j]
			}

		}

		if sum == instr.Target {
			return true
		}

	}

	return false

}

func partOne(instr []instruction) (result int) {

	for _, v := range instr {

		if isTruePartOne(v) {
			result += v.Target
		}

	}

	return

}

func base3String(n int, length int) (result string) {

	digits := make([]string, length)

	for i := 0; i < length; i++ {
		digits[i] = []string{"0", "1", "2"}[n%3]
		n /= 3
	}

	slices.Reverse(digits)
	result = strings.Join(digits, "")

	return

}

func concat(a int, b int) int {
	var pow int = 1
	for b >= pow {
		pow *= 10
	}
	return (a * pow) + b
}

func partTwo(instr []instruction) (result int) {

	for _, v := range instr {

		var possPermut int = int(math.Pow(3.0, float64(len(v.Components)-1)))

		for i := 0; i < possPermut; i++ {

			sum := v.Components[0]
			base3 := base3String(int(i), len(v.Components)-1)

			for j := 1; j < len(v.Components); j++ {

				operator := base3[j-1]

				if operator == '0' {
					sum *= v.Components[j]
				} else if operator == '1' {
					sum += v.Components[j]
				} else if operator == '2' {

					// Concat

					sumStr := strconv.Itoa(sum)
					numStr := strconv.Itoa(v.Components[j])

					sum, _ = strconv.Atoi(sumStr + numStr)

				}

			}

			if sum == v.Target {
				result += v.Target
				break
			}

		}

	}

	return
}
func main() {

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	parsed := parseData(data)

	fmt.Println(base3String(5, 4))

	fmt.Println(partOne(parsed))
	fmt.Println(partTwo(parsed))

}
