package main

import (
	"sync"
)

var num_decade map[int]int = make(map[int]int)
var mutex sync.Mutex
var producer_done bool = false

func main() {
	file_path := "./test_fajlovi/data_3000000.tsv"
	var decades *List = nil
	msgs := make(chan Person)
	done := make(chan int)
	// done_combiner := make(chan int)
	done_printer := make(chan int)
	p := &Producer{&msgs, &done}
	c := &Consumer{&msgs}
	c1 := &Consumer{&msgs}

	// combiner := &Combiner{decades, &done_combiner}
	printer := &Printer{decades, &done_printer, &msgs}

	go p.produce(file_path)
	go c.consume(100)
	go c1.consume(100)
	go printer.print(5)
	<-done_printer
}
