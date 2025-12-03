package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getHighestJoltage(amount int, nums []string, joltage string) string {
	if amount == 0 {
		return joltage
	}
	highestNum := "0"
	highestIdx := 0
	for i, num := range nums[:len(nums)-amount+1] {
		if num > highestNum {
			highestNum = num
			highestIdx = i
		}
	}
	return getHighestJoltage(amount-1, nums[highestIdx+1:], joltage+highestNum)
}

func main() {
	f, _ := os.Open("test_input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	part1 := 0
	part2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, "")
		joltagePart1, _ := strconv.Atoi(getHighestJoltage(2, nums, ""))
		joltagePart2, _ := strconv.Atoi(getHighestJoltage(12, nums, ""))
		part1 += joltagePart1
		part2 += joltagePart2
	}
	fmt.Printf("Part1: %d\n", part1)
	fmt.Printf("Part2: %d\n", part2)
}
