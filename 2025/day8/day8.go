package main

import (
	"bufio"
	"fmt"
	"maps"
	"math"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Vector struct {
	X, Y, Z float64
}

type VectorPair struct {
	First, Second Vector
}

func main() {
	f, _ := os.Open("test_input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	part1 := 0
	part2 := 0

	junctionBoxes := []Vector{}

	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		v := strings.Split(line, ",")
		x, _ := strconv.ParseFloat(v[0], 32)
		y, _ := strconv.ParseFloat(v[1], 32)
		z, _ := strconv.ParseFloat(v[2], 32)
		junction := Vector{X: x, Y: y, Z: z}
		junctionBoxes = append(junctionBoxes, junction)
	}

	distances := map[VectorPair]float64{}
	for i, v1 := range junctionBoxes {
		for _, v2 := range junctionBoxes[i+1:] {
			if v1 == v2 {
				continue
			}
			xSqrd := math.Pow(v1.X-v2.X, 2)
			ySqrd := math.Pow(v1.Y-v2.Y, 2)
			zSqrd := math.Pow(v1.Z-v2.Z, 2)
			linearDistance := math.Sqrt(xSqrd + ySqrd + zSqrd)
			distances[VectorPair{First: v1, Second: v2}] = linearDistance
		}
	}

	pairs := slices.Collect(maps.Keys(distances))
	sort.Slice(pairs, func(i, j int) bool {
		return distances[pairs[i]] < distances[pairs[j]]
	})

	circuits := map[Vector]int{}
	circuitTotals := map[int]int{}
	totalCircuits := 0
	for i, vectorPair := range pairs {
		firstCircuit, firstUsed := circuits[vectorPair.First]
		secondCircuit, secondUsed := circuits[vectorPair.Second]
		if firstUsed && secondUsed {
			if firstCircuit == secondCircuit {
				// fmt.Println("Nothing happens!")
			} else {
				circuitTotals[totalCircuits] = circuitTotals[firstCircuit] + circuitTotals[secondCircuit]
				delete(circuitTotals, firstCircuit)
				delete(circuitTotals, secondCircuit)
				for vec, c := range circuits {
					if c == firstCircuit || c == secondCircuit {
						circuits[vec] = totalCircuits
					}
				}
				totalCircuits++
				// fmt.Println("Merged!")
			}
		} else if secondUsed {
			circuits[vectorPair.First] = secondCircuit
			circuitTotals[secondCircuit]++
		} else if firstUsed {
			circuits[vectorPair.Second] = firstCircuit
			circuitTotals[firstCircuit]++
		} else {
			circuits[vectorPair.First] = totalCircuits
			circuits[vectorPair.Second] = totalCircuits
			circuitTotals[totalCircuits] = 2
			totalCircuits++
		}

		if i+1 == 1000 {
			totalsSorted := slices.Collect(maps.Values(circuitTotals))
			slices.Sort(totalsSorted)
			slices.Reverse(totalsSorted)
			part1 = totalsSorted[0] * totalsSorted[1] * totalsSorted[2]
		}

		if circuitTotals[totalCircuits-1] == len(junctionBoxes) {
			part2 = int(vectorPair.First.X * vectorPair.Second.X)
			break
		}
	}

	fmt.Printf("Part1: %d\n", part1)
	fmt.Printf("Part2: %d\n", part2)
}
