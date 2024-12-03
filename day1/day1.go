package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

func quickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	// Pick a pivot
	pivotIndex := len(a)-2 

	// Move the pivot to the right
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]

	// Go down the rabbit hole
	quickSort(a[:left])
	quickSort(a[left+1:])

	return a
}

func calculateTotalDistance(leftList []int, rightList []int) int {
	leftListSorted := quickSort(leftList)
	rightListSorted := quickSort(rightList)
	totalDistance := 0

	for i := 0; i < len(leftListSorted); i++ {
		distance := leftListSorted[i] - rightListSorted[i]
		if distance < 0 {
			totalDistance -= distance
		} else {
			totalDistance += distance
		}
	}

	return totalDistance
}


func calculateSimilarity(leftList []int, rightList []int) int { 
   rightMap := make(map[int]int) 
   similarity := 0

   for _, num := range rightList {
       _, ok := rightMap[num]
       if ok {
           rightMap[num]++;
       } else {
           rightMap[num] = 1 
       }
   }

   for _, num := range leftList {
       count, ok := rightMap[num]
       if ok {
         similarity += num * count
       }
   }

   return similarity
}

func main() {
	f, _ := os.Open("test_input.txt")
	defer f.Close()

	r := csv.NewReader(f)

	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	leftList := []int{}
	rightList := []int{}

	for _, record := range records {
		leftValue, err := strconv.Atoi(record[0])
		if err != nil {
			panic(err)
		}
		rightValue, err := strconv.Atoi(record[1])
		if err != nil {
			panic(err)
		}

		leftList = append(leftList, leftValue)
		rightList = append(rightList, rightValue)
	}

    println("Distance: ", calculateTotalDistance(leftList, rightList))
    println("Similiarity: ", calculateSimilarity(leftList, rightList))
}
