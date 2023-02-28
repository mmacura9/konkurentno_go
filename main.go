package main

import (
	"sync"
)

var num_decade map[int]int = make(map[int]int)
var mutex sync.Mutex
var producer_done bool = false

func main() {
	file_path := "./test_fajlovi/data.tsv"
	var decades *List = nil
	msgs := make(chan Person, 100)
	done := make(chan int, 1)
	// done_combiner := make(chan int)
	done_printer := make(chan int, 1)
	p := &Producer{&msgs, &done}
	c := &Consumer{&msgs}
	c1 := &Consumer{&msgs}
	c2 := &Consumer{&msgs}
	c3 := &Consumer{&msgs}
	c4 := &Consumer{&msgs}
	c5 := &Consumer{&msgs}

	// combiner := &Combiner{decades, &done_combiner}
	printer := &Printer{decades, &done_printer, &msgs}

	go p.produce(file_path)
	go c.consume(100)
	go c1.consume(100)
	go c2.consume(100)
	go c3.consume(100)
	go c4.consume(100)
	go c5.consume(100)

	go printer.print(5)
	<-done_printer
}
