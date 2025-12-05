package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	a int
	b int
}

func main() {
	f, _ := os.Open("test_input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	part1 := 0
	part2 := 0

	ranges := [](Range){}

	rangesParsed := false
	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		if line == "" {
			rangesParsed = true
			continue
		}
		if !rangesParsed {
			idRange := strings.Split(line, "-")
			lowEnd, _ := strconv.Atoi(idRange[0])
			highEnd, _ := strconv.Atoi(idRange[1])
			ranges = append(ranges, Range{a: lowEnd, b: highEnd})
		} else {
			id, _ := strconv.Atoi(line)
			for _, idRange := range ranges {
				if id >= idRange.a && id <= idRange.b {
					part1++
					break
				}
			}
		}
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].a < ranges[j].a
	})
	currentMax := 0
	for _, idRange := range ranges {
		lowEnd := max(currentMax+1, idRange.a)
		valid := idRange.b - lowEnd + 1
		if valid <= 0 {
			continue
		}
		part2 += valid
		currentMax = idRange.b
	}

	fmt.Printf("Part1: %d\n", part1)
	fmt.Printf("Part2: %d\n", part2)
}
