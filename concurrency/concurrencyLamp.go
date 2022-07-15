package concurrency

import "fmt"

type ConcurrencyLamp struct {
}

func NewConcurrencyLamp() *ConcurrencyLamp {
	return &ConcurrencyLamp{}
}

func (c *ConcurrencyLamp) TurnOn() {
	fmt.Println("Concurrency turn on.")
}

func (c *ConcurrencyLamp) TurnOff() {
	fmt.Println("Concurrency turn off.")
}
