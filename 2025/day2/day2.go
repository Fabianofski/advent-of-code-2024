package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func hasDoubleDigits(number string) bool {
	noOfDigits := len(number)
	half := noOfDigits / 2
	left := number[:half]
	right := number[half:]
	return left == right
}

func hasDoubleDigitsPart2(number string) bool {
	noOfDigits := len(number)
	for n := range noOfDigits / 2 {
		digits := n + 1
		if noOfDigits%digits != 0 {
			continue
		}
		pattern := true
		for i := 0; i+digits*2 <= noOfDigits; i += digits {
			leftNum := number[i : i+digits]
			rightNum := number[i+digits : i+digits*2]
			if leftNum != rightNum {
				pattern = false
			}
		}
		if pattern {
			return pattern
		}
	}
	return false
}

func main() {
	f, _ := os.Open("test_input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		val := strings.Split(line, ",")

		part1Score := 0
		part2Score := 0
		for _, y := range val {
			ranges := strings.Split(y, "-")
			left, _ := strconv.Atoi(ranges[0])
			right, _ := strconv.Atoi(ranges[1])
			for num := left; num <= right; num++ {
				if hasDoubleDigits(strconv.Itoa(num)) {
					part1Score += num
				}
				if hasDoubleDigitsPart2(strconv.Itoa(num)) {
					part2Score += num
				}
			}
		}
		fmt.Printf("Part1: %d\n", part1Score)
		fmt.Printf("Part2: %d\n", part2Score)
	}
}
