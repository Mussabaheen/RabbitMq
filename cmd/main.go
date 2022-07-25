package main

import (
	"fmt"
	"sync"

	"github.com/mussabaheen/rabbitmqgolang/src/Queues"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	go func(sync.WaitGroup) {
		Queues.Send(wg)
	}(wg)
	go func(sync.WaitGroup) {
		go Queues.Recieve(wg)
	}(wg)

	wg.Wait()
	fmt.Println("Go Routines Finished")
}
