package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Position struct {
	X, Y int
}

func main() {
	f, _ := os.Open("test_input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	splitters := map[Position]bool{}
	beams := map[Position]int{}
	start := Position{}
	height := 0

	part1 := 0
	part2 := 0

	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		row := strings.Split(line, "")
		for x, char := range row {
			if char == "." {
				continue
			} else if char == "S" {
				start = Position{X: x, Y: y}
				beams[start] = 1
			} else if char == "^" {
				splitters[Position{X: x, Y: y}] = true
			}
		}
		height++
	}

	for {
		newBeams := map[Position]int{}
		endReached := true
		for pos, timelines := range beams {
			if pos.Y >= height {
				continue
			}
			newPos := Position{X: pos.X, Y: pos.Y + 1}
			_, splitter := splitters[newPos]
			if splitter {
				leftBeam := Position{X: newPos.X + 1, Y: newPos.Y}
				rightBeam := Position{X: newPos.X - 1, Y: newPos.Y}
				newBeams[leftBeam] += timelines
				newBeams[rightBeam] += timelines
				endReached = false
				part1++
			} else {
				newBeams[newPos] += timelines
				endReached = false
			}
		}
		if endReached {
			break
		}
		beams = newBeams
	}

	for _, timelines := range beams {
		part2 += timelines
	}

	fmt.Printf("Part1: %d\n", part1)
	fmt.Printf("Part2: %d\n", part2)
}
