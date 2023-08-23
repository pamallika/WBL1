package main

import (
	"fmt"
	"sync"
)

func main() {
	//Используем WaitGroup чтобы дождаться отработки всех горутин, иначе выходим из программы досрочно
	var wg sync.WaitGroup
	data := [5]int{2, 4, 6, 8, 10}
	wg.Add(len(data))
	for i := range data {
		i := i
		go func() {
			fmt.Println(data[i] * data[i])
			wg.Done()
		}()
	}
	wg.Wait()
}
