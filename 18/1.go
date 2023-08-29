package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Counter1 struct {
	count int64
}

//Реализация через atomic(самый быстрый и подходящий способ реализации т.к. нужно выполнить только одну операцию)

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	var counter Counter1
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter.count, 1)
		}()
	}
	wg.Wait()
	fmt.Println(counter.count)
	fmt.Printf("With mutex: %f", time.Now().Sub(start).Seconds())
}
