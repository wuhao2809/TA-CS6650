package main

import (
	"fmt"
	"sync"
)

// QUESTIONS:
// 1. What will the final value of `counter` be after the program exits?
// 2. Are the FIRST THREE printed lines deterministic? If not, list what *must* be true

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	counter := 0

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 3; j++ {
				mu.Lock()
				counter++
				fmt.Printf("Goroutine %d incremented counter to %d\n", id, counter)
				mu.Unlock()
			}
		}(i)
	}

	wg.Wait()
	fmt.Println("Final counter:", counter)
}
