package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calcCheckSum(fileBlocks []string, moveWholeBlocks bool) int {
	for i := len(fileBlocks) - 1; i >= 0; i-- {
		end := fileBlocks[i]
		if end == "." {
			continue
		}

		blockLength := 1
		if moveWholeBlocks {
			blockLength = 0
			for j := i; j >= 0; j-- {
				next := fileBlocks[j]
				if next != end {
					break
				} else {
					blockLength++
				}
			}
		}

		for j := 0; j < len(fileBlocks); j++ {
			if j > i {
				i -= blockLength - 1
				break
			}

			freeSpace := 0
			for l := j; l < len(fileBlocks); l++ {
				next := fileBlocks[l]
				if next != "." {
					break
				} else {
					freeSpace++
				}
			}

			if freeSpace < blockLength {
				continue
			}

			for l := range blockLength {
				fileBlocks[j+l] = end
				fileBlocks[i-l] = "."
			}
			i -= blockLength - 1
			break
		}
	}

	checkSum := 0
	for i, block := range fileBlocks {
		if block == "." {
			continue
		}
		num, _ := strconv.Atoi(block)
		checkSum += i * num
	}

	return checkSum
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
	buffer := scanner.Text()
	numbers := strings.Split(buffer, "")
	id := 0

	fileBlocks := []string{}
	for i, num := range numbers {
		numConv, _ := strconv.Atoi(num)
		for range numConv {
			if i%2 == 0 {
				fileBlocks = append(fileBlocks, strconv.Itoa(id))
			} else {
				fileBlocks = append(fileBlocks, ".")
			}
		}
		if i%2 == 0 {
			id++
		}
	}

	countPartOne := calcCheckSum(append([]string{}, fileBlocks...), false)
	countPartTwo := calcCheckSum(append([]string{}, fileBlocks...), true)

	fmt.Printf("Part One: %d \n", countPartOne)
	fmt.Printf("Part Two: %d \n", countPartTwo)
}
