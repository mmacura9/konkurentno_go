package main

import (
	"sync"
)

func main() {
	var num_decade map[int]int = make(map[int]int)
	var mutex sync.Mutex
	var mutex1 sync.RWMutex
	var producer_done bool = false
	var consumers [5]*Consumer
	var wg sync.WaitGroup

	file_path := "./test_fajlovi/data.tsv"
	msgs := make(chan Person, 100)
	done_printer := make(chan int)

	p := new_Producer(msgs, &producer_done)
	printer := new_Printer(done_printer, msgs)
	N := 5

	go p.produce(file_path, &mutex1)

	for i := 0; i < N; i++ {
		consumers[i] = new_Consumer(msgs)
		wg.Add(1)
		go consumers[i].consume(100, &producer_done, num_decade, &wg, &mutex, &mutex1)
	}
	go printer.print(5, &producer_done, num_decade, &wg, &mutex, &mutex1)
	<-done_printer
}
