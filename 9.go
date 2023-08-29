package main

import (
	"fmt"
	"sync"
)

func write(numbers chan<- int, num int) {
	numbers <- num
}

func double(doubleNumbers chan int, num int, wg *sync.WaitGroup) {
	doubleNumbers <- num * 2
	wg.Done()
}

func main() {
	dataSize := 10
	numbers := make(chan int, dataSize)
	//Набиваем канал данными
	go func() {
		for i := 1; i <= dataSize; i++ {
			write(numbers, i)
		}
		close(numbers)
	}()
	doubleNumbers := make(chan int, dataSize)

	//Ждём отработки горутин, чтобы закрыть канал и вывести обработанные данные
	var wg sync.WaitGroup
	for num := range numbers {
		wg.Add(1)
		go double(doubleNumbers, num, &wg)
	}
	wg.Wait()
	close(doubleNumbers)

	for double := range doubleNumbers {
		fmt.Println(double)
	}
}
