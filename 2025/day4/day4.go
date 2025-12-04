package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanMap(cafeteria [][]string) [][]string {
	for y := range cafeteria {
		for x, item := range cafeteria[y] {
			if item == "X" {
				cafeteria[y][x] = "."
			}
		}
	}
	fmt.Println("\nCleaned:")
	printMap(cafeteria)
	return cafeteria
}

func printMap(cafeteria [][]string) {
	for y := range cafeteria {
		for _, item := range cafeteria[y] {
			fmt.Print(item)
		}
		fmt.Print("\n")
	}
}

func getRemoveableRolls(cafeteria [][]string) ([][]string, int) {
	fmt.Println("\nBefore::")
	printMap(cafeteria)
	rolls := 0
	for y, row := range cafeteria {
		for x, item := range cafeteria[y] {
			if item != "@" {
				continue
			}
			noOfRolls := 0
			for rollX := max(0, x-1); rollX <= min(x+1, len(row)-1); rollX++ {
				for rollY := max(0, y-1); rollY <= min(y+1, len(cafeteria)-1); rollY++ {
					if rollY == y && rollX == x {
						continue
					}
					if cafeteria[rollY][rollX] == "@" || cafeteria[rollY][rollX] == "X" {
						noOfRolls++
					}
				}
			}
			if noOfRolls <= 3 {
				cafeteria[y][x] = "X"
				rolls++
			}
		}
	}
	fmt.Println("\nRemoved:")
	printMap(cafeteria)
	return cafeteria, rolls
}

func main() {
	f, _ := os.Open("test_input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	cafeteria := [][]string{}
	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		cafeteria = append(cafeteria, []string{})
		for _, item := range strings.Split(line, "") {
			cafeteria[y] = append(cafeteria[y], item)
		}
	}
	cafeteria, part1 := getRemoveableRolls(cafeteria)
	cafeteria = cleanMap(cafeteria)
	part2 := part1
	for {
		cafeteria, rolls := getRemoveableRolls(cafeteria)
		cafeteria = cleanMap(cafeteria)
		if rolls == 0 {
			break
		}
		part2 += rolls
	}

	fmt.Printf("Part1: %d\n", part1)
	fmt.Printf("Part2: %d\n", part2)
}
