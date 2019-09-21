package main

import (
	"fmt"
	"sync"
	"time"
)

func main()  {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("1st goroutine sleeping...")
		time.Sleep(1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2nd goroutine sleeping...")
		time.Sleep(1)
	}()

	wg.Wait()
	fmt.Println("complete")

	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("Hello %v!\n", id)
	}

	const numGreeters = 5
	wg.Add(numGreeters)
	for i:= 0; i < numGreeters; i++ {
		go hello(&wg, i+1)
	}
	wg.Wait()
}

