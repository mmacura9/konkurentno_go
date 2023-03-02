package main

import (
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
	defer wg.Done()
	break_for := false
	for !break_for {
		select {
		case person := <-*c.msgs:
			do(&num, N, decade_map, &person)
		default:
			if get_producer_done() {
				mutex.Lock()
				for key, val := range decade_map {
					num_decade[key] += val
				}
				mutex.Unlock()
				break_for = true
			}
		}

	}
	// fmt.Println("Consumer done")
}
