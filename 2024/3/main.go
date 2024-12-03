package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const validStart = "mul("
const validEnd = ')'

const doMultiply = "do()"
const dontMultiply = "don't()"

const validChars = "1234567890,"

func parseData(data []byte) (result [][]int) {

	result = make([][]int, 0)
	stringResult := make([]string, 0)

	dataString := string(data)

	var inInstruction = false

	var currentInstruction = ""

	for i := 0; i < len(dataString); i++ {

		if i+len(validStart) > len(dataString) && !inInstruction {
			break
		} else if !inInstruction && dataString[i:i+len(validStart)] == validStart {

			inInstruction = true
			i += len(validStart) - 1

		} else if inInstruction {

			if strings.Contains(validChars, string(dataString[i])) {
				currentInstruction += string(dataString[i])
			} else if !strings.Contains(validChars, string(dataString[i])) && dataString[i] != validEnd {

				i = (i - len(currentInstruction))
				currentInstruction = ""
				inInstruction = false

				continue

			} else if dataString[i] == validEnd {

				inInstruction = false

				stringResult = append(stringResult, currentInstruction)
				currentInstruction = ""

			}

		}

	}

	for _, s := range stringResult {

		nums := strings.Split(s, ",")
		num1, err := strconv.Atoi(nums[0])
		if err != nil {
			panic(err)
		}

		num2, err := strconv.Atoi(nums[1])
		if err != nil {
			panic(err)
		}

		result = append(result, []int{num1, num2})

	}

	return

}

func partOne(data [][]int) (total int) {

	for _, v := range data {

		total += v[0] * v[1]

	}

	return

}

func partTwo(data []byte) (total int) {

	dataString := string(data)

	var inInstruction = false
	var multiply = true

	var currentInstruction = ""

	for i := 0; i < len(dataString); i++ {

		if i+len(doMultiply) < len(dataString) && dataString[i:i+len(doMultiply)] == doMultiply {
			multiply = true
		} else if i+len(dontMultiply) < len(dataString) && dataString[i:i+len(dontMultiply)] == dontMultiply {
			multiply = false
		}

		if i+len(validStart) > len(dataString) && !inInstruction {
			break
		} else if !inInstruction && dataString[i:i+len(validStart)] == validStart {

			inInstruction = true
			i += len(validStart) - 1

		} else if inInstruction {

			if strings.Contains(validChars, string(dataString[i])) {
				currentInstruction += string(dataString[i])
			} else if !strings.Contains(validChars, string(dataString[i])) && dataString[i] != validEnd {

				i = (i - len(currentInstruction))
				currentInstruction = ""
				inInstruction = false

				continue

			} else if dataString[i] == validEnd {

				inInstruction = false

				nums := strings.Split(currentInstruction, ",")
				num1, err := strconv.Atoi(nums[0])
				if err != nil {
					panic(err)
				}

				num2, err := strconv.Atoi(nums[1])
				if err != nil {
					panic(err)
				}

				if multiply {

					total += num1 * num2

				}

				currentInstruction = ""

			}

		}

	}

	return

}

func main() {
	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Println(partOne(parseData(bytes)))
	fmt.Println(partTwo(bytes))
}
