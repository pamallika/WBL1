package main

import (
	"fmt"
	"sync"
	"time"
)

// Реализайция через mutex (медленнее чем через atomic, но позволяет в будущем выполнить несколько операций в рамках лока одного mutex)
type Counter struct {
	sync.Mutex
	count int
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	var counter Counter
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer counter.Unlock()
			counter.Lock()
			counter.count++
		}()
	}
	wg.Wait()
	fmt.Println(counter.count)
	fmt.Printf("With mutex: %f", time.Now().Sub(start).Seconds())
}
