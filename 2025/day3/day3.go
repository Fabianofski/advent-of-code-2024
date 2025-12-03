package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, _ := os.Open("test_input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	part1 := 0
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, "")

		numConcat := "0"
		for i, num1 := range nums {
			for _, num2 := range nums[i+1:] {
				concat := num1 + num2
				if concat > numConcat {
					numConcat = concat
				}
			}
		}
		joltage, _ := strconv.Atoi(numConcat)
		part1 += joltage
	}
	fmt.Printf("Part1: %d\n", part1)
}
