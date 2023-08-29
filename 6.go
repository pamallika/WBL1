package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Отслеживая канал при записи в который будет происходить выход(в канал можно передавать и другие типы, но
// пустая структура весит меньше всех и не будет выделаять память, копировать элементы и т.д.)
func gorutine1(exit chan struct{}) {
	for {
		select {
		case <-exit:
			fmt.Println("gorutine1 stopped")
			return
		}
	}
}

// через закрытие канала
func gorutine2(ch chan bool) {
	for {
		// Проверяем доступность канала чтобы не пытаться читать из закрытого(panic)
		_, open := <-ch
		if !open {
			fmt.Println("gorutine1 stopped")
			return
		}
	}
}

// через sync.WaitGroup
func gorutine3(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("gorutine3 stopped")
}

// через context cancel()
func gorutine4(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("gorutine4 stopped")
			return
		}
	}
}

// через context timeout
func gorutine5(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("gorutine5 stopped")
			return
		}
	}
}

func main() {
	//через пустую структуру. аналогично можно через bool / любой тип но пустая структура весит практически ничего
	ch := make(chan struct{})
	fmt.Println("gorutine1 starts")
	go gorutine1(ch)
	ch <- struct{}{}
	close(ch)

	fmt.Println("gorutine2 starts")
	ch1 := make(chan bool)
	go gorutine2(ch1)
	close(ch1)

	//через sync.WaitGroup
	var wg sync.WaitGroup
	wg.Add(1)
	fmt.Println("gorutine3 starts")
	go gorutine3(&wg)
	wg.Wait()

	//через контекст
	ctxB := context.Background()
	ctx, cancel := context.WithCancel(ctxB)
	fmt.Println("gorutine4 starts")
	go gorutine4(ctx)
	cancel()

	//через контекст  с таймаутом
	ctx, _ = context.WithTimeout(ctxB, 3*time.Second)
	fmt.Println("gorutine5 starts")
	go gorutine5(ctx)

	select {
	case <-time.After(time.Second * 3):
		fmt.Println("everything stopped")
	}
}
