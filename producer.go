package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sync"
)

type Producer struct {
	msgs chan Person
	done *bool
}

func new_Producer(msgs chan Person, done *bool) *Producer {
	return &Producer{msgs, done}
}

func (p *Producer) produce(file_path string, mutex *sync.RWMutex) {
	csvFile, err := os.Open(file_path)

	if err != nil {
		fmt.Println(err)
	}

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)

	reader.Comma = '\t'
	reader.LazyQuotes = true
	reader.FieldsPerRecord = -1

	csvLines, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {
		if len(line) != 6 {
			continue
		}
		person := new_Person(line[0], line[1], line[2], line[3], line[4], line[5])
		p.msgs <- person
	}
	mutex.Lock()
	*p.done = true
	mutex.Unlock()
}
