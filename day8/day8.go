package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Vector2 struct {
	X, Y int
}

func (v1 Vector2) Add(v2 Vector2) Vector2 {
	return Vector2{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
	}
}

func (v1 Vector2) Subtract(v2 Vector2) Vector2 {
	return Vector2{
		X: v1.X - v2.X,
		Y: v1.Y - v2.Y,
	}
}

func isInBounds(spot Vector2, mapSize Vector2) bool {
	aboveZero := spot.X >= 0 && spot.Y >= 0
	belowMapLength := spot.X < mapSize.X && spot.Y < mapSize.Y
	return aboveZero && belowMapLength
}

func getAnodeLocations(spotA Vector2, spotB Vector2, mapSize Vector2, infinite bool) []Vector2 {
	anodeLocations := []Vector2{}
	distance := spotB.Subtract(spotA)

	anodeLocations = append(anodeLocations, spotA)
	anodeLocations = append(anodeLocations, spotB)
	for {
		spotA = spotA.Subtract(distance)
		if !isInBounds(spotA, mapSize) {
			break
		}
		anodeLocations = append(anodeLocations, spotA)
		if !infinite {
			break
		}
	}

	for {
		spotB = spotB.Add(distance)
		if !isInBounds(spotB, mapSize) {
			break
		}
		anodeLocations = append(anodeLocations, spotB)
		if !infinite {
			break
		}
	}
	return anodeLocations
}

func printMap(antennas map[string][]Vector2, anodes map[string][]Vector2, mapSize Vector2) {
	fmt.Println()
	for y := range mapSize.Y {
		for x := range mapSize.X {
			pos := Vector2{X: x, Y: y}
			obj := "."
			for char := range antennas {
				if slices.Contains(antennas[char], pos) {
					obj = char
					break
				}
			}
			if obj == "." {
				for char := range antennas {
					if slices.Contains(anodes[char], pos) {
						obj = "#"
						break
					}
				}
			}
			fmt.Print(obj)
		}
		fmt.Println()
	}
}

func calcDistinctAnodes(antennas map[string][]Vector2, anodes map[string][]Vector2, mapSize Vector2, partTwo bool) int {
	for char, spots := range antennas {
		for i, spotA := range spots {
			for _, spotB := range spots[i+1:] {
				locations := getAnodeLocations(spotA, spotB, mapSize, partTwo)
				anodes[char] = append(anodes[char], locations...)
			}
		}
	}

	distinctAnodes := map[Vector2]bool{}
	for char, spots := range anodes {
		for _, spot := range spots {
			isAntenna := false
			for _, antenna := range antennas[char] {
				if antenna == spot {
					isAntenna = true
					break
				}
			}
			if !isAntenna || partTwo {
				distinctAnodes[spot] = true
			}
		}

	}
	printMap(antennas, anodes, mapSize)

	return len(distinctAnodes)
}

func main() {
	f, err := os.Open("test_input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	mapSize := Vector2{X: 0, Y: 0}
	antennas := map[string][]Vector2{}
	anodes := map[string][]Vector2{}

	scanner := bufio.NewScanner(f)
	for y := 0; scanner.Scan(); y++ {
		mapSize.Y++
		line := scanner.Text()
		chars := strings.Split(line, "")
		mapSize.X = len(chars)
		for x, char := range chars {
			if char == "." {
				continue
			}
			antennas[char] = append(antennas[char], Vector2{X: x, Y: y})
		}
	}

	fmt.Printf("Part One: %d \n", calcDistinctAnodes(antennas, anodes, mapSize, false))
	fmt.Printf("Part Two: %d \n", calcDistinctAnodes(antennas, anodes, mapSize, true))
}
