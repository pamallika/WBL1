package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func publisher(channel chan<- int, exit chan struct{}) {
	for {
		select {
		case <-exit:
			return
		default:
			channel <- rand.Intn(9)
		}
	}
}

func producer(channel <-chan int, exit chan struct{}) {
	for {
		select {
		case <-exit:
			return
		default:
			fmt.Fprintf(os.Stdout, "%d\n", <-channel)
		}
	}
}

func main() {
	var timeC time.Duration
	fmt.Println("Введите время работы(секунды):")
	fmt.Scanf("%d", &timeC)

	exit := make(chan struct{})
	channel := make(chan int)

	//Сначала нужно завершить цикл, а потом закрывать каналы, чтобы не было паники при записи и чтении из закрытого
	defer close(channel)
	defer close(exit)
	//записываем и читаем из канала
	go publisher(channel, exit)
	go producer(channel, exit)
	//после истечения заданного времени, кидаем структуру в канал, функция её отловит и выйдет из цикла
	select {
	case <-time.After(timeC * time.Second):
		exit <- struct{}{}
	}
}
