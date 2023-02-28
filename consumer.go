package main

import (
	"strconv"
	"strings"
)

type Consumer struct {
	msgs *chan Person
}

func (c *Consumer) consume(N int) {
	num := 0
	decade_map := make(map[int]int)

	for {
		if producer_done && len(*c.msgs) == 0 {
			break
		}
		person := <-*c.msgs
		num++
		if num == N {
			num = 0
			mutex.Lock()
			for key, val := range decade_map {
				num_decade[key] += val
				decade_map[key] = 0
			}
			mutex.Unlock()
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

}
