package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	calculationsEnabled bool = true
)

func check_for_valid_mul_part1(line string) int {
	if line[:4] != "mul(" {
		return 0
	}

	params := ""
	for j := 4; ; j++ {
		char := string(line[j])
		if char == ")" {
			values := strings.Split(params, ",")
			if len(values) < 2 {
				return 0
			}
			firstValInt, _ := strconv.Atoi(values[0])
			secondValInt, _ := strconv.Atoi(values[1])
			return firstValInt * secondValInt
		} else {
			params += char
		}
	}

}

func check_for_valid_mul_part2(line string) int {
	if line[:4] == "do()" {
		calculationsEnabled = true
	}
	if line[:7] == "don't()" {
		calculationsEnabled = false
	}

	if !calculationsEnabled {
		return 0
	}

	return check_for_valid_mul_part1(line)
}

func main() {
	f, err := os.Open("test_input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	sum_part1 := 0
	sum_part2 := 0
	for scanner.Scan() {
		line := scanner.Text()
		for i := range line[:len(line)-8] {
			sum_part1 += check_for_valid_mul_part1(line[i:])
			sum_part2 += check_for_valid_mul_part2(line[i:])
		}
	}
	fmt.Println("Sum Part 1:", sum_part1)
	fmt.Println("Sum Part 2:", sum_part2)
}
