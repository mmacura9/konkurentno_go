package main

import (
	"fmt"
	"sort"
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

		time.Sleep(2 * time.Second)
	}
	fmt.Println("Final: ")
	wg.Wait()
	keys := make([]int, 0, len(num_decade))
	for k := range num_decade {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(p, q int) bool {
		return keys[p] < keys[q]
	})

	for _, k := range keys {
		fmt.Println(k, "-", k+9, "\t", num_decade[k])
	}

	fmt.Println("Done forever")
	*p.done <- 1
}
