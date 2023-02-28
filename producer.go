package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Producer struct {
	msgs *chan Person
	done *chan int
}

func (p *Producer) produce(file_path string) {
	csvFile, err := os.Open(file_path)

	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	reader.Comma = '\t'
	reader.LazyQuotes = true

	csvLines, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
	}

	for _, line := range csvLines {
		person := Person{line[0], line[1], line[2], line[3], line[4], line[5]}
		*p.msgs <- person
	}
	producer_done = true
	*p.done <- 1
}
