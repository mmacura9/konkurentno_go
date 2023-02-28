package main

import (
	"fmt"
	"time"
)

type Printer struct {
	decades *List
	done    *chan int
	msgs    *chan Person
}

func (p *Printer) print(N int) {
	for !producer_done || len(*p.msgs) > 0 {
		fmt.Println("Start")

		for key, val := range num_decade {
			fmt.Println(key, "-", key+9, "\t", val)
		}
		// d1 := p.decades
		// for d1 != nil {
		// 	fmt.Println(*d1.start_decade, "-", *d1.start_decade+9, " ", *d1.num_of_alive)
		// 	d1 = d1.next
		// }
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Final: ")
	for key, val := range num_decade {
		fmt.Println(key, "-", key+9, "\t", val)
	}
	// d1 := p.decades
	// for d1 != nil {
	// 	fmt.Println(*d1.start_decade, "-", *d1.start_decade+9, " ", *d1.num_of_alive)
	// 	d1 = d1.next
	// }
	*p.done <- 1
}
