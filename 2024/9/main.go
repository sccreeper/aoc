package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"time"
)

var blankArray []int

func init() {

	blankArray = make([]int, 1000)

	for i := 0; i < 1000; i++ {
		blankArray[i] = -1
	}

}

func genNumArray(val int, length int) (result []int) {
	result = make([]int, 0, length)

	for i := 0; i < length; i++ {
		result = append(result, val)
	}

	return
}

func parseData(data []byte) (result []int, files int) {

	dataString := string(data)

	result = make([]int, 0)

	var freeSpace bool = false
	var fileIdCount int

	for _, v := range dataString {

		num, err := strconv.Atoi(string(v))
		if err != nil {
			panic(err)
		}

		if freeSpace {
			result = append(result, genNumArray(-1, num)...)
		} else {
			result = append(result, genNumArray(fileIdCount, num)...)
			if num != 0 {
				fileIdCount++
			}
		}

		freeSpace = !freeSpace

	}

	for i := 0; i < len(result); i++ {
		if result[i] != -1 {
			files++
		}
	}

	return

}

func isSorted(diskMap []int, fileCount int) (sorted bool) {
	sorted = true

	for k := 0; k < fileCount; k++ {
		if diskMap[k] == -1 {
			sorted = false
			break
		}
	}

	return

}

func partOne(diskMap []int, fileCount int) (checksum int) {

	for i := len(diskMap) - 1; i >= 0; i-- {
		if isSorted(diskMap, fileCount) {
			break
		}

		if diskMap[i] == -1 {
			continue
		} else {

			// Find earliest -1
			var freeSpaceStart int
			for j := 0; j < len(diskMap); j++ {
				if diskMap[j] == -1 {
					freeSpaceStart = j
					break
				}
			}

			emptyBlockSize := 0

			for l := freeSpaceStart; l < len(diskMap); l++ {
				if diskMap[l] == -1 {
					emptyBlockSize++
				} else {
					break
				}
			}

			copySize := 0

			for l := i; l >= 0 && copySize < emptyBlockSize && diskMap[l] != -1; l-- {
				copySize++
			}

			blockToCopy := diskMap[i-copySize+1 : i+1]
			slices.Reverse(blockToCopy)

			copy(diskMap[freeSpaceStart:freeSpaceStart+copySize], blockToCopy)
			copy(diskMap[i-copySize+1:i+1], blankArray[:copySize])

			if isSorted(diskMap, fileCount) {
				break
			}

		}
	}

	// Calculate checksum

	for i := 0; i < len(diskMap) && diskMap[i] != -1; i++ {

		checksum += diskMap[i] * i

	}

	return

}

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	startTime := time.Now().UnixMicro()

	fmt.Println("Parsing data...")
	parsed, filesCount := parseData(data)
	fmt.Println("Part one...")
	fmt.Printf("%d\n", partOne(parsed, filesCount))

	fmt.Printf("Finished in %dus\n", time.Now().UnixMicro()-startTime)
}
