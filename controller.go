package main

type Controller struct {
	consumers_done []chan int
}

func new_Controller() *Controller {
	consumers_done := make([]chan int, 0)
	return &Controller{consumers_done: consumers_done}
}

func (c *Controller) add_consumer(con chan int) {
	c.consumers_done = append(c.consumers_done, con)
}

func (c *Controller) control(producer *Producer, msgs chan Person) {
	<-producer.done
	msgs_empty := make(chan int)

	go func() {
		for len(msgs) != 0 {
		}
		msgs_empty <- 1
	}()
	<-msgs_empty
	for _, channel := range c.consumers_done {
		channel <- 1
	}
}
