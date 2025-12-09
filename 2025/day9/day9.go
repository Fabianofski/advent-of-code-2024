package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Position struct {
	X, Y int
}

type Area struct {
	Tile1, Tile2 Position
	Area         int
}

func calcArea(pos1 Position, pos2 Position) int {
	width := int(math.Abs(float64(pos1.X-pos2.X))) + 1
	height := int(math.Abs(float64(pos1.Y-pos2.Y))) + 1
	area := width * height
	return area
}

func containsAir(pos1 Position, pos2 Position, bounds map[int]Position) bool {
	bl := Position{X: min(pos1.X, pos2.X), Y: min(pos1.Y, pos2.Y)}
	tr := Position{X: max(pos1.X, pos2.X), Y: max(pos1.Y, pos2.Y)}
	for y := bl.Y; y <= tr.Y; y++ {
		if bounds, exists := bounds[y]; exists {
			if bl.X < bounds.X || tr.X > bounds.Y {
				return true
			}
		} else {
			return true
		}
	}
	return false
}

func calcLargestArea(redTiles []Position, tiles map[int]Position) (Area, Area) {
	largestArea1 := Area{}
	largestArea2 := Area{}
	for i, tile1 := range redTiles {
		for _, tile2 := range redTiles[i+1:] {
			area := calcArea(tile1, tile2)
			if area > largestArea1.Area {
				largestArea1.Area = area
				largestArea1.Tile1 = tile1
				largestArea1.Tile2 = tile2
			}

			if area > largestArea2.Area && !containsAir(tile1, tile2, tiles) {
				largestArea2.Area = area
				largestArea2.Tile1 = tile1
				largestArea2.Tile2 = tile2
			}
		}
	}
	return largestArea1, largestArea2
}

func getDirection(pos1 Position, pos2 Position) Position {
	direction := Position{X: 0, Y: 0}
	xDiff := pos2.X - pos1.X
	yDiff := pos2.Y - pos1.Y
	if xDiff < 0 {
		direction.X = -1
	} else if xDiff > 0 {
		direction.X = 1
	}
	if yDiff < 0 {
		direction.Y = -1
	} else if yDiff > 0 {
		direction.Y = 1
	}

	return direction
}

func fillGreenTiles(redTiles []Position) map[int]Position {
	boundaries := map[int]Position{}
	for i, tile := range redTiles {
		nextTile := redTiles[(i+1)%len(redTiles)]
		direction := getDirection(tile, nextTile)
		newTile := Position{X: tile.X, Y: tile.Y}
		for {
			if bounds, exists := boundaries[newTile.Y]; exists {
				if newTile.X > bounds.Y {
					bounds.Y = newTile.X
				}
				if newTile.X < bounds.X {
					bounds.X = newTile.X
				}
				boundaries[newTile.Y] = bounds
			} else {
				boundaries[newTile.Y] = Position{X: newTile.X, Y: newTile.X}
			}
			if newTile == nextTile {
				break
			}
			newTile.X += direction.X
			newTile.Y += direction.Y
		}
	}

	return boundaries
}

func main() {
	f, _ := os.Open("test_input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	redTiles := []Position{}
	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		v := strings.Split(line, ",")
		x, _ := strconv.Atoi(v[0])
		y, _ := strconv.Atoi(v[1])
		tile := Position{X: x, Y: y}
		redTiles = append(redTiles, tile)
	}

	bounds := fillGreenTiles(redTiles)
	area1, area2 := calcLargestArea(redTiles, bounds)
	fmt.Printf("Part1: %d\n", area1.Area)
	fmt.Printf("Part2: %d\n", area2.Area)
}
