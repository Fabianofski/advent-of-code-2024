package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)


func main() {
	f, err := os.Open("test_input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
    rules := map[string][]string{}
    updates := [][]string{} 

	for scanner.Scan() {
		line := scanner.Text()
        if strings.Contains(line, "|"){
            values := strings.Split(line, "|")
            rules[values[0]] = append(rules[values[0]], values[1])
        } else if line != "" {
            updates = append(updates, strings.Split(line, ","))
        }
	}

    count_part_one := 0
    count_part_two := 0

    for _, update := range updates {
        valid := true
        for i := 0; i < len(update); i++ {
            number := update[i]
            for _, rule := range rules[number] {
                previous_nums := update[:i]
                index := slices.Index(previous_nums, rule)
                if index != -1 {
                    update = slices.Delete(update, index, index+1)
                    update = slices.Insert(update, i, rule)
                    valid = false
                    i = 0
                    break
                }
            } 
        }
        val, _ := strconv.Atoi(update[len(update)/2])
        if valid {
            count_part_one += val
        } else {
            count_part_two += val
        }
    }

    fmt.Printf("Part One: %d \n", count_part_one)
    fmt.Printf("Part Two: %d \n", count_part_two)
}
