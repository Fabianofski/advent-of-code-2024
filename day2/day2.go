package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(value int) int {
	if value < 0 {
		return -value
	} else {
		return value
	}
}

func list_is_safe(values []int) bool {
	fmt.Println(values)
	if len(values) == 1 {
		return true
	}
	increasing := values[len(values)-1] > values[0]

	for i, value := range values[:len(values)-1] {
		difference := values[i+1] - value

		if increasing && difference < 0 {
			fmt.Println("Increasing List but now decreasing")
			return false
		} else if !increasing && difference > 0 {
			fmt.Println("Decreasing List but now increasing")
			return false
		} else if difference == 0 {
			fmt.Println("Difference is 0")
			return false
		} else if abs(difference) > 3 {
			fmt.Printf("Difference %d greater than 3\n", difference)
			return false
		}
	}
	return true
}

func list_is_safe_with_dampener(values []int) bool {
	if list_is_safe(values) {
		return true
	} else {
		for i := 0; i < len(values); i++ {
			temp := make([]int, 0, len(values)-1)
			temp = append(temp, values[:i]...)
			temp = append(temp, values[i+1:]...)
			if list_is_safe(temp) {
				return true
			}
		}
	}
	return false
}

func main() {
	f, err := os.Open("test_input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	safeCount := 0
	safeDampenedCount := 0
	unsafeCount := 0
	unsafeDampenedCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		var values []int

		for _, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				fmt.Println("Error converting to integer:", err)
				return
			}
			values = append(values, num)

		}
		fmt.Println(values)
		if list_is_safe(values) {
			fmt.Println("Safe")
			safeCount++
		} else {
			fmt.Println("Unsafe")
			unsafeCount++
		}

		if list_is_safe_with_dampener(values) {
			fmt.Println("Safe Dampened")
			safeDampenedCount++
		} else {
			fmt.Println("Unsafe Dampened")
			unsafeDampenedCount++
		}

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Printf("Part1: Found %d safe lists and %d unsafe lists out of %d lists\n", safeCount, unsafeCount, safeCount+unsafeCount)
	fmt.Printf("Part2: Found %d safe lists and %d unsafe lists out of %d lists\n", safeDampenedCount, unsafeDampenedCount, safeDampenedCount+unsafeDampenedCount)
}
