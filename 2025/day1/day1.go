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

	for scanner.Scan() {
		line := scanner.Text()
		negative := line[0] == 'L'
		value, err := strconv.Atoi(line[1:])
		if err != nil {
			panic("???")
		}

		fmt.Printf("Left: %t by %d\n", negative, value)
		if negative {
			score = (score - value) % 100
			if score < 0 {
				score = 100 + score
			}
		} else {
			score = (score + value) % 100
		}
		fmt.Printf("Score: %d\n", score)

		if score == 0 {
			code++
		}
	}
	println(code)
}
