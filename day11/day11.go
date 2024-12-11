package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calcNumberOfDigits(i int) int {
	if i == 0 {
		return 1
	}
	count := 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
}

func blink(stones map[int]int, max int) int {
	newStones := map[int]int{}
	alreadyCalculated := map[int][]int{}
	for range max {
		newStones = map[int]int{} 
		for stone, no := range stones {
			val, ok := alreadyCalculated[stone]
			if ok {
                for _, v := range val {
                    newStones[v] += no 
                }
			} else if stone == 0 {
                newStones[1] += no
				alreadyCalculated[stone] = []int{1}
			} else {
				digits := calcNumberOfDigits(stone)
				if digits%2 == 0 {
					stoneStr := strconv.Itoa(stone)
					firstStone, _ := strconv.Atoi(stoneStr[0 : digits/2])
					secondStone, _ := strconv.Atoi(stoneStr[digits/2:])

                    newStones[firstStone] += no
                    newStones[secondStone] += no
					alreadyCalculated[stone] = []int{firstStone, secondStone}
				} else {
                    res := stone * 2024
                    newStones[res] += no
					alreadyCalculated[stone] = []int{res}
				}
			}
		}
		stones = newStones
	}

    noOfStones := 0
    for _, no := range newStones {
        noOfStones += no
    }

	return noOfStones
}

func main() {
	f, err := os.Open("test_input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Scan()

	stones := map[int]int{}
	line := scanner.Text()
	split := strings.Split(line, " ")
	for _, stone := range split {
		stoneNum, _ := strconv.Atoi(stone)
        stones[stoneNum]++
	}

    partOne := blink(stones, 25)
	partTwo := blink(stones, 75)

	fmt.Printf("Part One: %d \n", partOne)
	fmt.Printf("Part Two: %d \n", partTwo)
}
