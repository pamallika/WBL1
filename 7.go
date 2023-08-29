package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func writeMap(data map[int]int, id int, m *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	//Блокируем доступ к мапе во время записи
	m.Lock()
	data[id] = id * id
	m.Unlock()
}

func main() {
	var mutex sync.Mutex
	var wg sync.WaitGroup

	n := rand.Intn(10)
	mapa := make(map[int]int)
	for i := 1; i < n+1; i++ {
		wg.Add(1)
		go writeMap(mapa, i, &mutex, &wg)
	}
	wg.Wait()
	fmt.Println(mapa)

}
