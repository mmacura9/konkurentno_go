package main

import (
	"fmt"
	"sync"
	"time"
)

type Printer struct {
	done *chan int
	msgs *chan Person
}

func (p *Printer) print(N int, wg *sync.WaitGroup) {
	for !get_producer_done() || len(*p.msgs) > 0 {
		fmt.Println("Start")
		mutex.Lock()
		for key, val := range num_decade {
			fmt.Println(key, "-", key+9, "\t", val)
		}
		mutex.Unlock()
		fmt.Println("End")

		time.Sleep(1 * time.Second)
	}
	fmt.Println("Final: ")
	wg.Wait()

	for key, val := range num_decade {
		fmt.Println(key, "-", key+9, "\t", val)
	}

	fmt.Println("Done forever")
	*p.done <- 1
}
