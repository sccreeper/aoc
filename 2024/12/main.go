package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"math/rand"
	"os"
	"time"
)

type Flower struct {
	Adjacent []*Flower
	Value    byte
	Position [2]int
	InArea   bool
}

var lookVectors [][2]int = [][2]int{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

func parseData(data []byte) (flowers [][]*Flower) {

	var x int
	var y int

	result2d := make([][]*Flower, 0)
	accumulator := make([]*Flower, 0)

	for _, b := range data {

		if b == '\n' {
			x = 0
			y++
			result2d = append(result2d, accumulator)
			accumulator = make([]*Flower, 0)
			continue
		} else {

			accumulator = append(accumulator, &Flower{
				Value:    b,
				Position: [2]int{x, y},
				Adjacent: make([]*Flower, 0),
				InArea:   false,
			})

			x++
		}

	}

	result2d = append(result2d, accumulator)

	// Add adjacent flowers (only if of same type)

	for y := 0; y < len(result2d); y++ {
		for x := 0; x < len(result2d[0]); x++ {

			for _, v := range lookVectors {

				if x+v[0] < 0 || y+v[1] < 0 || x+v[0] >= len(result2d[0]) || y+v[1] >= len(result2d) {
					continue
				} else if result2d[y+v[1]][x+v[0]].Value == result2d[y][x].Value {

					if !containsFlower(result2d[y+v[1]][x+v[0]], result2d[y][x].Adjacent) {
						result2d[y][x].Adjacent = append(result2d[y][x].Adjacent, result2d[y+v[1]][x+v[0]])
					}

				}

			}

		}
	}

	flowers = result2d

	return

}

const emptySpace = '.'

func containsFlower(f *Flower, flowers []*Flower) bool {

	for _, v := range flowers {
		if v.Position[0] == f.Position[0] && v.Position[1] == f.Position[1] && v.Value == f.Value {
			return true
		}
	}

	return false

}

func fill(start *Flower, areaValues []*Flower) (result []*Flower) {
	result = make([]*Flower, 0)

	start.InArea = true

	for _, v := range start.Adjacent {

		if v.InArea {
			continue
		} else {
			areaValues = fill(v, areaValues)
		}

	}

	areaValues = append(areaValues, start)

	return areaValues

}

func visualise(areas [][]*Flower, width int, height int) {

	vImg := image.NewNRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})

	r := rand.New(rand.NewSource(time.Now().Unix()))

	for _, a := range areas {

		areaColourR := r.Intn(255)
		areaColourG := r.Intn(255)
		areaColourB := r.Intn(255)

		for _, v := range a {
			vImg.Set(v.Position[0], v.Position[1], color.RGBA{uint8(areaColourR), uint8(areaColourG), uint8(areaColourB), 255})
		}

	}

	f, err := os.Create("out.png")
	if err != nil {
		panic(err)
	}

	err = png.Encode(f, vImg)
	if err != nil {
		panic(err)
	}

}

func partOne(flowers [][]*Flower) (total int) {

	// Flatten flowers

	flattenedFlowers := make([]*Flower, 0)

	for y := 0; y < len(flowers); y++ {
		for x := 0; x < len(flowers[0]); x++ {

			flattenedFlowers = append(flattenedFlowers, flowers[y][x])

		}
	}

	// Find areas

	areas := make([][]*Flower, 0)

	for i := 0; i < len(flattenedFlowers); i++ {

		if flattenedFlowers[i].InArea {
			continue
		} else {
			areas = append(areas, fill(flattenedFlowers[i], make([]*Flower, 0)))
		}

	}

	// Calculate perimeters

	for _, a := range areas {

		var perimeterMap map[[2]int]int = make(map[[2]int]int)

		for _, f := range a {

			for _, v := range lookVectors {

				lookCoords := [2]int{f.Position[0] + v[0], f.Position[1] + v[1]}

				if lookCoords[0] < 0 || lookCoords[1] < 0 || lookCoords[0] >= len(flowers[0]) || lookCoords[1] >= len(flowers) {

					perimeterMap[lookCoords]++

				} else if flowers[lookCoords[1]][lookCoords[0]].Value != a[0].Value {

					perimeterMap[lookCoords]++

				}

			}

		}

		var perimeterLength int

		for _, v := range perimeterMap {
			perimeterLength += v
		}

		total += len(a) * perimeterLength

	}

	visualise(areas, len(flowers[0]), len(flowers))

	return

}

func partTwo(flowers [][]*Flower) (total int) {
	return
}

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	parsed := parseData(data)
	fmt.Println(partOne(parsed))
}
