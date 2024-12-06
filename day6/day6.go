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

func checkGuardsPath(gameField [][]string, pos Vector2, dir Vector2, checkLoops bool) (int, int, bool) {
	visited := map[Vector2][]Vector2{}
	obstacles := []Vector2{}
	startPos := pos
	startDir := dir

	for {
		newPos := pos.Add(dir)

		directions, _ := visited[newPos]
		if slices.Contains(directions, dir) {
			return len(visited), len(obstacles), true
		}

		outOfBoundsY := newPos.Y >= len(gameField) || newPos.Y < 0
		outOfBoundsX := newPos.X >= len(gameField) || newPos.X < 0
		if outOfBoundsX || outOfBoundsY {
			break
		}

		obj := gameField[newPos.Y][newPos.X]
		if obj == "." {
			newPosIsStartPos := startPos.X == newPos.X && startPos.Y == newPos.Y
			if checkLoops && !newPosIsStartPos && !slices.Contains(obstacles, newPos) {
				gameField[newPos.Y][newPos.X] = "O"
				_, _, isLoop := checkGuardsPath(gameField, startPos, startDir, false)
				gameField[newPos.Y][newPos.X] = obj
				if isLoop {
					obstacles = append(obstacles, newPos)
				}
			}

			pos = newPos
			visited[pos] = append(visited[pos], dir)
		} else {
			dir = Vector2{X: dir.Y * -1, Y: dir.X}
		}
	}

	return len(visited), len(obstacles), false
}

func main() {
	f, err := os.Open("test_input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	gameField := [][]string{}
	pos := Vector2{}
	dir := Vector2{X: 0, Y: -1}

	for i := 0; scanner.Scan(); i++ {
		row := scanner.Text()
		game_row := []string{}
		for j, obj := range strings.Split(row, "") {
			if obj == "^" {
				pos.X = j
				pos.Y = i
				obj = "."
			}
			game_row = append(game_row, obj)
		}
		gameField = append(gameField, game_row)
	}

	tileCount, loopCount, _ := checkGuardsPath(gameField, pos, dir, true)

	fmt.Printf("Part One: %d \n", tileCount)
	fmt.Printf("Part Two: %d \n", loopCount)
}
