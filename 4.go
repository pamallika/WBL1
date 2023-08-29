package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

func worker(i uint, queue <-chan int) {
	for d := range queue {
		//получаем число и пишем в поток
		fmt.Fprintf(os.Stdout, "Worker: %d Value: %d %s\n", i, d)
	}
}

// воркеры завершаются при завершении работы main, после приема сигнала прерывания и остановки посыла данных в канал
func main() {
	//считываем N воркеров
	var workersCount uint
	var i uint
	fmt.Println("Введите количество воркеров: ")
	_, err := fmt.Scanf("%d", &workersCount)
	if err != nil {
		log.Fatalf("Input error count workers: %s", err)
	}
	//канал для передачи данных
	queue := make(chan int)
	//канал ожидающий прерывания
	exit := make(chan os.Signal, 1)

	for i = 1; i <= workersCount; i++ {
		//запускаем воркеров
		go worker(i, queue)

	}
	//Отлавливаем прерывание
	signal.Notify(exit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	//в бесконечном цикле ждем, когда поступит сигнал прерывания, в случае если поступил, закрываем канал и возвращаемся из main
	for {
		select {
		//обработка прерывания
		case <-exit:
			close(queue)
			fmt.Fprintln(os.Stdout, "Workers stopped")
			return
		default:
			queue <- rand.Intn(9)
		}
	}
}
