package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	c := make(chan interface{})
	go func() {
		time.Sleep(5*time.Second)
		close(c)
	}()

	fmt.Println("Blocking in read...")
	select {
	case <-c:
		fmt.Printf("Unblocking %v later.\n", time.Since(start))
	}
}
