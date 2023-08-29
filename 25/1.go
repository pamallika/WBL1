package main

import (
	"context"
	"fmt"
	"time"
)

// sleepCtx через context
func sleepCtx(c context.Context, t time.Duration) {
	ctx, _ := context.WithTimeout(c, t)
	select {
	case <-ctx.Done():
		return
	}
}

func main() {
	fmt.Println("Write time to sleep(seconds): ")
	var timeout int
	fmt.Scan(&timeout)
	ctx := context.Background()
	//задаем время в секундах
	duration := time.Duration(timeout) * time.Second
	fmt.Printf("Start waiting\n")
	sleepCtx(ctx, duration)
	fmt.Printf("Waited for %f seconds\n", duration.Seconds())
}
