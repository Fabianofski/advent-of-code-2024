package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func solveEquation(result int, numbers []int, operations int) bool {
	length := len(numbers) - 1
    combinations := int(math.Pow(float64(operations), float64(length)))

	for i := 0; i < combinations; i++ {
		sum := numbers[0]

		for j := length - 1; j >= 0; j-- {
			operation :=  (i / int(math.Pow(float64(operations), float64(j)))) % operations 
			num := numbers[length-j]

			if operation == 0 {
				sum *= num
			} else if operation == 1 {
				sum += num
			} else {
                sum = concatenateNums(sum, num)
            }

			if sum > result {
				break
			}
		}
		if sum == result {
			return true
		}
	}
	return false
}

func concatenateNums(a int, b int) int {
	strA := strconv.Itoa(a)
	strB := strconv.Itoa(b)

	concatenated := strA + strB

	result, _ := strconv.Atoi(concatenated)
	return result
}

func main() {
	f, err := os.Open("test_input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	countPartOne := 0
	countPartTwo := 0

	scanner := bufio.NewScanner(f)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		equation := strings.Split(line, " ")

		result, _ := strconv.Atoi(strings.Replace(equation[0], ":", "", 1))
		numbers := []int{}
        for _, num := range equation[1:] {
			conv, _ := strconv.Atoi(num)
			numbers = append(numbers, conv)
		}

		if solveEquation(result, numbers, 2) {
			countPartOne += result
		}
		if solveEquation(result, numbers, 3) {
			countPartTwo += result
		}
	}

	fmt.Printf("Part One: %d \n", countPartOne)
	fmt.Printf("Part Two: %d \n", countPartTwo)
}
