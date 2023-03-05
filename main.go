package main

import (
	"sync"
)

func main() {
	var num_decade map[int]int = make(map[int]int)
	var mutex sync.Mutex
	var mutex1 sync.Mutex
	var wg sync.WaitGroup

	consumers := make([]*Consumer, 0)
	producer_done := false
	p_done := make(chan int)

	file_path := "./test_fajlovi/data.tsv"
	msgs := make(chan Person, 100)
	done_printer := make(chan int)
	controller := new_Controller()

	producer := new_Producer(msgs, p_done, &producer_done)
	printer := new_Printer(done_printer, msgs)
	N := 10

	go controller.control(producer, msgs)
	go producer.produce(file_path, &mutex1)

	for i := 0; i < N; i++ {
		consumers = append(consumers, new_Consumer(msgs, controller))
		wg.Add(1)
		go consumers[i].consume(100, num_decade, &wg, &mutex)
	}

	go printer.print(5, &producer_done, num_decade, &wg, &mutex, &mutex1)
	<-done_printer
}
