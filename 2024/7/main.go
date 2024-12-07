package main

import (
	"fmt"
	"io"
	"math"
	"os"
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

func partOne(instr []instruction) (result int) {

	for _, v := range instr {

		var possPermut uint64 = uint64(math.Pow(2.0, float64(len(v.Components)-1)))

		var i uint64
		for i = 0; i < possPermut; i++ {

			sum := v.Components[0]

			for j := 1; j < len(v.Components); j++ {

				operator := (i >> (j - 1)) & 1

				if operator == 1 {
					sum *= v.Components[j]
				} else {
					sum += v.Components[j]
				}

			}

			if sum == v.Target {
				result += sum
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

	fmt.Println(partOne(parseData(data)))

}
