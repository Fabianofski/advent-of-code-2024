package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("test_input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	score := 50
	code := 0
	codePart2 := 0
	fmt.Printf("Score: %d\n", score)

	for scanner.Scan() {
		line := scanner.Text()
		negative := line[0] == 'L'
		value, err := strconv.Atoi(line[1:])
		if err != nil {
			panic("???")
		}

		if negative {
			fmt.Printf("Left by %d\n", value)
		} else {
			fmt.Printf("Right by %d\n", value)
		}

		distanceToZero := 0
		if negative {
			if score == 0 {
				distanceToZero = 100
			} else {
				distanceToZero = score
			}
			score = (score - value) % 100
			if score < 0 {
				score = 100 + score
			}
		} else {
			if score == 0 {
				distanceToZero = 100
			} else {
				distanceToZero = 100 - score
			}
			score = (score + value) % 100
		}

		if value >= distanceToZero {
			codePart2++
			value -= distanceToZero
		}

		for {
			value -= 100
			if value < 0 {
				break
			}
			codePart2++
		}

		if score == 0 {
			code++
		}
		fmt.Printf("Score: %d\n", score)
		fmt.Println()
	}
	fmt.Printf("Code Part 1 %d\n", code)
	fmt.Printf("Code Part 2 %d\n", codePart2)
}
