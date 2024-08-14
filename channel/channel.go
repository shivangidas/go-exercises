package main

import (
	"context"
	"fmt"
	"time"
)

func odd(ch chan<- int, ctx context.Context) {
	defer close(ch)
	for i := 1; i < 20; i += 2 {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			return
		case ch <- i:
		}
	}
}
func even(ch chan<- int) {
	defer close(ch)
	for i := 2; i <= 20; i += 2 {
		time.Sleep(2 * time.Second)
		ch <- i
	}
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		odd(ch1, ctx)
	}()
	go func() {
		even(ch2)
	}()
	for i := range ch1 {
		// val, ok := <-ch1
		// if ok {
		// 	fmt.Println(val)
		// }
		if i > 10 {
			break
		}
		fmt.Println(i)
		fmt.Println(<-ch2)
	}
}
