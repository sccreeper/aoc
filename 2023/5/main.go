package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

const (
	SeedToSoil         = "seed-to-soil"
	SoilToFertilizer   = "soil-to-fertilizer"
	FertilizerToWater  = "fertilizer-to-water"
	WaterToLight       = "water-to-light"
	LightToTemp        = "light-to-temperature"
	TempToHumidity     = "temperature-to-humidity"
	HumidityToLocation = "humidity-to-location"
)

var stages []string = []string{
	SeedToSoil,
	SoilToFertilizer,
	FertilizerToWater,
	WaterToLight,
	LightToTemp,
	TempToHumidity,
	HumidityToLocation,
}

const integers = "1234567890"

type LocationMap struct {
	SourceStart      int
	DestinationStart int
	Range            int
}

type Almanac struct {
	Seeds []int

	Maps map[string][]LocationMap
}

func get_numbers(line string) []int {

	numbers := make([]int, 0)

	for _, v := range strings.Split(line, " ") {

		if v == "" {
			continue
		} else {

			i, err := strconv.Atoi(v)
			if err != nil {
				fmt.Printf("Cannot convert %s to integer", v)
				panic(err)
			}

			numbers = append(numbers, i)

		}

	}

	return numbers

}

func get_location_map(numbers []int) LocationMap {

	return LocationMap{
		SourceStart:      numbers[1],
		DestinationStart: numbers[0],
		Range:            numbers[2],
	}

}

func parse_file(data []byte) Almanac {

	al := Almanac{}
	al.Maps = make(map[string][]LocationMap)
	lines := strings.Split(string(data), "\n")

	al.Seeds = get_numbers(strings.Split(lines[0], ":")[1])

	current_data := SeedToSoil

	// Parse first line (seeds)
	//3 (line 4) is the first line relevant data appears on

	for i := 3; i < len(lines); i++ {

		if lines[i] == "" {
			continue
		}

		if strings.Contains(lines[i], ":") {
			current_data = strings.Split(lines[i], " ")[0]
			continue
		}

		switch current_data {
		case SeedToSoil:
			al.Maps[SeedToSoil] = append(al.Maps[SeedToSoil], get_location_map(get_numbers(lines[i])))
		case SoilToFertilizer:
			al.Maps[SoilToFertilizer] = append(al.Maps[SoilToFertilizer], get_location_map(get_numbers(lines[i])))
		case FertilizerToWater:
			al.Maps[FertilizerToWater] = append(al.Maps[FertilizerToWater], get_location_map(get_numbers(lines[i])))
		case WaterToLight:
			al.Maps[WaterToLight] = append(al.Maps[WaterToLight], get_location_map(get_numbers(lines[i])))
		case LightToTemp:
			al.Maps[LightToTemp] = append(al.Maps[LightToTemp], get_location_map(get_numbers(lines[i])))
		case TempToHumidity:
			al.Maps[TempToHumidity] = append(al.Maps[TempToHumidity], get_location_map(get_numbers(lines[i])))
		case HumidityToLocation:
			al.Maps[HumidityToLocation] = append(al.Maps[HumidityToLocation], get_location_map(get_numbers(lines[i])))
		}
	}

	return al

}

func part_1(al Almanac) int {

	var locations []int = al.Seeds

	for i := range al.Seeds {

		for _, stage := range stages {

			var destination int = locations[i]

			for _, location_map := range al.Maps[stage] {

				if destination <= location_map.SourceStart+location_map.Range-1 && destination >= location_map.SourceStart {
					//We're in the range
					destination = location_map.DestinationStart + (destination - location_map.SourceStart)
					break
				}

			}

			locations[i] = destination

		}

	}

	//fmt.Println(locations)

	slices.Sort(locations)

	return locations[0]

}

// Part 2 just calls part 1 but modifies the seed numbers
func part_2(al Almanac) int {

	seeds := make([]int, 0)

	for i := 0; i < len(al.Seeds); i += 2 {

		fmt.Printf("Seed range %d to %d\n", al.Seeds[i], al.Seeds[i]+al.Seeds[i+1]-1)

		//Generate range

		seed_range := make([]int, 0)

		for j := al.Seeds[i]; j < al.Seeds[i]+al.Seeds[i+1]; j++ {
			seed_range = append(seed_range, j)
		}

		seeds = append(
			seeds,
			part_1(Almanac{
				Seeds: seed_range,
				Maps:  al.Maps,
			},
			),
		)

	}

	fmt.Printf("Performing final sort on array length %d\n", len(seeds))

	slices.Sort(seeds)

	return seeds[0]
}

func main() {

	var file_name string
	fmt.Println("Enter the your file name")
	fmt.Scanln(&file_name)

	file_bytes, err := os.ReadFile(file_name)
	if err != nil {
		panic(err)
	}

	start := time.Now()

	almanac := parse_file(file_bytes)

	fmt.Println(part_2(almanac))

	fmt.Println(time.Since(start))

}
