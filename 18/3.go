package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter3 struct {
	count int
}

// Реализация через канал (самый медленный)
func main() {
	start := time.Now()
	var wg sync.WaitGroup
	var counter Counter3
	channel := make(chan struct{}, 1)
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//Блокируем запись, т.к. заняли канал
			channel <- struct{}{}
			counter.count++
			//освобождаем канал и открываем запись после проделанных манипуляций
			<-channel
		}()
	}
	wg.Wait()
	fmt.Println(counter.count)
	fmt.Printf("With channel: %f", time.Now().Sub(start).Seconds())

}
