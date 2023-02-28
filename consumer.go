package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type Consumer struct {
	msgs   *chan Person
	p_done *chan int
}

func newConsumer(msgs *chan Person) *Consumer {
	return &Consumer{msgs: msgs}
}

func do(num *int, N int, decade_map map[int]int, person *Person) {
	*num = *num + 1
	if *num == N {
		*num = 0
		mutex.Lock()
		for key, val := range decade_map {
			num_decade[key] += val
		}
		mutex.Unlock()

		for key, _ := range decade_map {
			decade_map[key] = 0
		}
	}
	if strings.Contains(person.primaryProfession, "actor") || strings.Contains(person.primaryProfession, "actress") {
		_, err := strconv.Atoi(person.deathYear)
		if err != nil {

			birth, err := strconv.Atoi(person.birthYear)

			if err == nil {
				decade := birth
				decade = decade / 10
				decade = decade * 10
				decade_map[decade]++
			}
		}
	}
}

func (c *Consumer) consume(N int, wg *sync.WaitGroup) {
	num := 0
	decade_map := make(map[int]int)

	for {
		if get_producer_done() && len(*c.msgs) == 0 {
			mutex.Lock()
			for key, val := range decade_map {
				num_decade[key] += val
			}
			mutex.Unlock()
			break
		}
		select {
		case <-*c.p_done:
			var person *Person = nil
			mutex2.Lock()
			if len(*c.msgs) != 0 {
				p1 := <-*c.msgs
				person = &p1
			}
			mutex2.Unlock()
			if person != nil {
				do(&num, N, decade_map, person)

			}
		case person := <-*c.msgs:
			do(&num, N, decade_map, &person)
		}

	}
	wg.Done()
	fmt.Println("Consumer done")
}
