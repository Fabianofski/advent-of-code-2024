package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calcResult(nums []int, operator string) int {
	result := nums[0]
	fmt.Print(result)
	for _, num := range nums[1:] {
		if operator == "*" {
			result *= num
			fmt.Print("*")
		} else {
			result += num
			fmt.Print("+")
		}
		fmt.Print(num)
	}
	fmt.Printf("=%d\n", result)
	return result
}

func main() {
	f, _ := os.Open("test_input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	part1 := 0
	part2 := 0

	lines := []string{}
	equations := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		row := strings.Fields(line)
		equations = append(equations, row)
	}

	fmt.Println("\n\nPart1: ")
	operators := equations[len(equations)-1]
	for i, operator := range operators {
		part1Nums := []int{}
		for _, row := range equations[:len(equations)-1] {
			strValue := row[i]
			value, _ := strconv.Atoi(strValue)
			part1Nums = append(part1Nums, value)
		}
		part1 += calcResult(part1Nums, operator)
	}

	fmt.Println("\n\nPart2: ")
	no := 0
	values := []int{}
	for i := range len(lines[0]) {
		number := ""
		for _, line := range lines[:len(lines)-1] {
			char := line[i]
			if char != ' ' {
				number += string(char)
			}
		}
		if number == "" {
			part2 += calcResult(values, operators[no])
			values = []int{}
			no++
		} else {
			value, _ := strconv.Atoi(number)
			values = append(values, value)
			if i+1 == len(lines[0]) {
				part2 += calcResult(values, operators[no])
			}
		}
	}
	fmt.Printf("Part1: %d\n", part1)
	fmt.Printf("Part2: %d\n", part2)
}
