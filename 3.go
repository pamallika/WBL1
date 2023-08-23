package main

import (
	"fmt"
	"time"
)

func main() {
	sum := 0
	data := [5]int{2, 4, 6, 8, 10}
	sumDouble := make(chan int, len(data))

	//Кидаем квадраты чисел из массива в канал конкурентно
	for i := range data {
		go doubleInt(sumDouble, data[i])
	}
	//Суммируем все числа из канала
	for i := 0; i < len(data); i++ {
		sum = sum + <-sumDouble
	}
	close(sumDouble)
	fmt.Println(sum)
}

func doubleInt(channel chan int, number int) {
	time.Sleep(3 * time.Second)
	channel <- number * number
	return
}
