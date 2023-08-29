package main

import (
	"fmt"
	"time"
)

func sleep(seconds int) {
	start := time.Now()
	for i, timeout := 0, time.After(time.Duration(seconds)*time.Second); ; {
		select {
		case <-timeout:
			return
		default:
			fmt.Println(time.Now().Sub(start).Seconds())
		}
		i++
	}
}

func main() {
	fmt.Println("Write time to sleep(seconds) : ")
	var timeout int
	fmt.Scan(&timeout)
	sleep(timeout)
}
