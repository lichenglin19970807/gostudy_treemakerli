// Package concurrency show my study of concurrency programming
package concurrency

import (
	"fmt"
	"sync"
)

type counter struct {
	CounterType int
	Name        string

	sync.Mutex
	count uint64
}

func (c *counter) getCount() uint64 {
	c.Lock()
	defer c.Unlock()
	return c.count
}

func (c *counter) Incr() {
	c.Lock()
	defer c.Unlock()
	c.count++
}

func MutexDemo() {
	var c counter
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				c.Incr()
			}
		}()
	}
	wg.Wait()
	fmt.Println(c.getCount())
}
