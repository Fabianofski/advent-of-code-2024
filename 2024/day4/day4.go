package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func check_for_xmas(line string, term string) int {
	return strings.Count(line, term) + strings.Count(line, reverse(term))
}

func check_part_one(input [][]string, term string) {
	count := 0

	// check horizontally
	for y := 0; y < len(input); y++ {
		line := ""
		for x := 0; x < len(input[0]); x++ {
			line += input[y][x]
		}
		count += check_for_xmas(line, term)
	}

	// check vertically
	for x := 0; x < len(input[0]); x++ {
		line := ""
		for y := 0; y < len(input); y++ {
			line += input[y][x]
		}
		count += check_for_xmas(line, term)
	}

	// check up diagonal
	for y := 0; y < len(input)*2; y++ {
		line := ""
		for x := 0; x < len(input[0])*2; x++ {
			if y-x >= 0 && y-x < len(input) && x < len(input[0]) {
				line += input[y-x][x]
			}
		}
		count += check_for_xmas(line, term)
	}

	// check down diagonal
	for y := -len(input); y < len(input); y++ {
		line := ""
		for x := 0; x < len(input[0])*2; x++ {
			if y+x >= 0 && y+x < len(input) && x < len(input[0]) {
				line += input[y+x][x]
			}
		}
		count += check_for_xmas(line, term)
	}

	fmt.Printf("Part One: Found XMAS %d times\n", count)
}

func check_part_two(input [][]string) {
	count := 0

	for y := 0; y < len(input)-2; y++ {
		for x := 0; x < len(input[y])-2; x++ {
			diagonal_up := ""
			diagonal_down := ""
			for i := 0; i <= 2; i++ {
				diagonal_up += input[y+i][x+i]
				diagonal_down += input[y+2-i][x+i]
			}
			if check_for_xmas(diagonal_down, "MAS") == 1 && check_for_xmas(diagonal_up, "MAS") == 1 {
				count++
			}
		}
	}

	fmt.Printf("Part Two: Found X-MAS %d times\n", count)
}

func main() {
	f, err := os.Open("test_input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	input := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, strings.Split(line, ""))
	}
	check_part_one(input, "XMAS")
	check_part_two(input)
}
