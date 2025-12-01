package main

import (
	"encoding/csv"
	"os"
)

func main() {
	f, _ := os.Open("test_input.txt")
	defer f.Close()

	r := csv.NewReader(f)

	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

	for _, record := range records {
		println(record)
	}
}
