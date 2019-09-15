package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	salutation := "hello"
	wg.Add(1)
	go func() {
		defer wg.Done()
		salutation = "welcome"
	}()

	wg.Wait()
	fmt.Println(salutation)

	for _, greeting := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(greeting string) {
			defer wg.Done()
			fmt.Println(greeting)
		}(greeting)
	}
	wg.Wait()
}

