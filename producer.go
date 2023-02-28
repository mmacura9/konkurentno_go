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

func set_producer_done() {
	mutex1.Lock()
	producer_done = true
	mutex1.Unlock()
}

func get_producer_done() bool {
	mutex1.RLock()
	x := producer_done
	mutex1.RUnlock()
	return x
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
	reader.FieldsPerRecord = -1

	csvLines, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
	}

	for _, line := range csvLines {
		if len(line) != 6 {
			continue
		}
		person := Person{line[0], line[1], line[2], line[3], line[4], line[5]}
		*p.msgs <- person
	}
	set_producer_done()
	for i := 0; i < 6; i++ {
		*p.done <- 1
	}

}
