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

func printMap(tiles []Position) {
	width := 0
	height := 0
	tilesMap := map[Position]string{}
	for _, tile := range tiles {
		tilesMap[tile] = "X"
		if tile.X > width {
			width = tile.X + 2
		}
		if tile.Y > height {
			height = tile.Y + 2
		}
	}
	for y := range height {
		for x := range width {
			if val, ok := tilesMap[Position{X: x, Y: y}]; ok {
				fmt.Print(val)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func calcArea(pos1 Position, pos2 Position) int {
	width := int(math.Abs(float64(pos1.X-pos2.X))) + 1
	height := int(math.Abs(float64(pos1.Y-pos2.Y))) + 1
	area := width * height
	return area
}

func isInBetween(largestArea Area, pos Position) bool {
	betweenX := (pos.X >= largestArea.Tile1.X && pos.X <= largestArea.Tile2.X) || (pos.X >= largestArea.Tile2.X && pos.X <= largestArea.Tile1.X)
	betweenY := (pos.Y >= largestArea.Tile1.Y && pos.Y <= largestArea.Tile2.Y) || (pos.Y >= largestArea.Tile2.Y && pos.Y <= largestArea.Tile1.Y)
	return betweenX && betweenY
}

func calcLargestArea(redTiles []Position, tiles []Position) Area {
	largestArea := Area{}
	for i, tile1 := range redTiles {
		for _, tile2 := range tiles[i+1:] {
			if largestArea.Area != 0 && isInBetween(largestArea, tile2) && isInBetween(largestArea, tile1) {
				continue
			}
			area := calcArea(tile1, tile2)
			if area > largestArea.Area {
				largestArea.Area = area
				largestArea.Tile1 = tile1
				largestArea.Tile2 = tile2
			}
		}
	}
	return largestArea
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

func fillGreenTiles(redTiles []Position) []Position {
	greenAndRedTiles := []Position{}
	for i, tile := range redTiles {
		nextTile := redTiles[(i+1)%len(redTiles)]
		direction := getDirection(tile, nextTile)
		newTile := Position{X: tile.X, Y: tile.Y}
		for {
			greenAndRedTiles = append(greenAndRedTiles, newTile)
			if newTile == nextTile {
				break
			}
			newTile.X += direction.X
			newTile.Y += direction.Y
		}
	}

	return greenAndRedTiles
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

	greenAndRedTiles := fillGreenTiles(redTiles)
	printMap(redTiles)

	part1 := calcLargestArea(redTiles, redTiles).Area
	fmt.Printf("Part1: %d\n", part1)

	part2 := calcLargestArea(redTiles, greenAndRedTiles).Area
	printMap(greenAndRedTiles)
	fmt.Printf("Part2: %d\n", part2)
}
