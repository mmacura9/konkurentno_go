package main

import (
	"sync"
)

var num_decade map[int]int = make(map[int]int)
var mutex sync.Mutex
var mutex2 sync.Mutex
var mutex1 sync.RWMutex

var producer_done bool = false

func main() {
	file_path := "./test_fajlovi/data.tsv"
	msgs := make(chan Person, 100)
	done := make(chan int)
	done_printer := make(chan int, 1)
	p := &Producer{&msgs, &done}
	printer := &Printer{&done_printer, &msgs}
	N := 5
	var consumers [5]*Consumer
	var wg sync.WaitGroup

	// combiner := &Combiner{decades, &done_combiner}
	go p.produce(file_path)
	for i := 0; i < N; i++ {
		consumers[i] = &Consumer{&msgs, &done}
		wg.Add(1)
		go consumers[i].consume(100, &wg)
	}

	go printer.print(5, &wg)
	<-done_printer
}
