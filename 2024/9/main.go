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

// Finds the biggest space that will fit a given file
func findFileSpace(diskMap []int, fileLength int) (spaceIndex int) {

	var inFileSpace bool = false
	var fileSpaceSize int = 0

	for i := 0; i < len(diskMap); i++ {

		if diskMap[i] == -1 && inFileSpace {
			fileSpaceSize++
		} else if diskMap[i] != -1 && inFileSpace {

			inFileSpace = false

			if fileLength <= fileSpaceSize {
				return i - fileSpaceSize
			} else {
				fileSpaceSize = 0
				continue
			}

		} else if diskMap[i] == -1 && !inFileSpace {
			inFileSpace = true
			fileSpaceSize++
		}
	}

	return -1

}

func prettifyDiskMap(diskMap []int) (result string) {

	result = ""

	for _, v := range diskMap {

		if v == -1 {
			result += "."
		} else {
			result += strconv.Itoa(v)
		}

	}

	return

}

func partTwo(diskMap []int) (checksum int) {

	// Map of file IDs, their leftmost indexes and their length
	var fileIDattr map[int][2]int = make(map[int][2]int)
	var fileIDs []int = make([]int, 0)

	var currentFile int = -1
	var currentFileLength int = 0
	var inFile bool = false

	for i := len(diskMap) - 1; i >= 0; i-- {

		if inFile {
			if diskMap[i] != currentFile {

				fileIDattr[currentFile] = [2]int{i + 1, currentFileLength}
				fileIDs = append(fileIDs, currentFile)

				i++
				currentFile = -1
				inFile = false
				currentFileLength = 0
			} else {
				currentFileLength++
			}
		} else {
			if diskMap[i] == -1 {
				continue
			} else {
				currentFile = diskMap[i]
				currentFileLength = 1
				inFile = true
			}
		}

	}

	// fmt.Println(fileIDattr)

	// Sort file IDs then loop through in descending order
	slices.Sort(fileIDs)

	for i := len(fileIDs) - 1; i >= 0; i-- {
		leftMostIndex := fileIDattr[fileIDs[i]][0]
		fileLength := fileIDattr[fileIDs[i]][1]

		fileIndex := findFileSpace(diskMap[:leftMostIndex], fileLength)

		if fileIndex != -1 {

			fileSource := diskMap[leftMostIndex : leftMostIndex+fileIndex]

			copy(diskMap[fileIndex:fileIndex+fileLength], fileSource)
			copy(diskMap[leftMostIndex:leftMostIndex+fileLength], blankArray[:fileLength])

			// fmt.Println(prettifyDiskMap(diskMap))

		}

	}

	// Calculate checksum

	var checksumEnd int
	for i := len(diskMap) - 1; i >= 0; i-- {
		if diskMap[i] != -1 {
			checksumEnd = i + 1
			break
		}
	}

	for i := 0; i < checksumEnd; i++ {

		if diskMap[i] == -1 {
			continue
		}

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
	// parsed, filesCount := parseData(data)
	// fmt.Println("Part one...")
	// fmt.Printf("%d\n", partOne(parsed, filesCount))
	parsed, _ := parseData(data)
	fmt.Println("Part two...")
	fmt.Printf("%d\n", partTwo(parsed))

	fmt.Printf("Finished in %dus\n", time.Now().UnixMicro()-startTime)
}
