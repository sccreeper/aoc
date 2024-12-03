package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type mulInstruction struct {
	DoMultiply bool
	Num1       int
	Num2       int
}

const validStart = "mul("
const validEnd = ')'

const doMultiply = "do()"
const dontMultiply = "don't()"

const validChars = "1234567890,"

func parseData(data []byte) (result []mulInstruction) {
	result = make([]mulInstruction, 0)

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

				result = append(result, mulInstruction{
					Num1:       num1,
					Num2:       num2,
					DoMultiply: multiply,
				})

				currentInstruction = ""

			}

		}

	}

	return

}

func partOne(data []mulInstruction) (total int) {

	for _, v := range data {

		total += v.Num1 * v.Num2

	}

	return

}

func partTwo(data []mulInstruction) (total int) {

	for _, v := range data {

		if v.DoMultiply {
			total += v.Num1 * v.Num2
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
	fmt.Println(partTwo(parseData(bytes)))
}
