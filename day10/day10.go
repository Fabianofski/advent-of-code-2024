package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

var (
	directions []Vector2 = []Vector2{
		{X: 1, Y: 0},
		{X: -1, Y: 0},
		{X: 0, Y: 1},
		{X: 0, Y: -1},
	}
	checkedPositions []Vector2 = []Vector2{}
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

func followTrail(trails [][]int, position Vector2, trailScore *int, onlyDistinct bool) {
	if !onlyDistinct && slices.Contains(checkedPositions, position) {
		return
	}
	checkedPositions = append(checkedPositions, position)
	currentElevation := trails[position.Y][position.X]
	if currentElevation == 9 {
		*trailScore++
		return
	}

	for _, direction := range directions {
		newPos := position.Add(direction)
		if newPos.X < 0 || newPos.Y < 0 || newPos.X >= len(trails[0]) || newPos.Y >= len(trails) {
			continue
		}
		newElevation := trails[newPos.Y][newPos.X]
		if newElevation-currentElevation == 1 {
			followTrail(trails, newPos, trailScore, onlyDistinct)
		}
	}
}

func main() {
	f, err := os.Open("test_input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	trails := [][]int{}
	trailHeads := []Vector2{}

	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		elevations := strings.Split(line, "")
		elevationsNum := []int{}
		for x, elevation := range elevations {
			elevationNum, _ := strconv.Atoi(elevation)
			elevationsNum = append(elevationsNum, elevationNum)
			if elevationNum == 0 {
				trailHeads = append(trailHeads, Vector2{X: x, Y: y})
			}
		}
		trails = append(trails, elevationsNum)
	}

	countPartOne := 0
	countPartTwo := 0
	for _, head := range trailHeads {
		trailScore := 0
		checkedPositions = []Vector2{}
		followTrail(trails, head, &trailScore, false)
		countPartOne += trailScore

		trailScore = 0
		followTrail(trails, head, &trailScore, true)
		countPartTwo += trailScore
	}

	fmt.Printf("Part One: %d \n", countPartOne)
	fmt.Printf("Part Two: %d \n", countPartTwo)
}
