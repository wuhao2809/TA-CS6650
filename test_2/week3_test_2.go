package main

import (
	"fmt"
	"sync"
)

// QUESTION:
// 1. What is gonna be printed on the terminal? Why?

type KV struct {
	mu      sync.Mutex
	value   string
	version int
}

// put updates the value only if the version matches expectedVersion
func (kv *KV) put(newValue string, expectedVersion int) bool {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	if kv.version != expectedVersion {
		return false // reject update
	}
	kv.value = newValue
	kv.version++
	return true
}

func main() {
	kv := &KV{value: "init", version: 0}
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		ok_1 := kv.put("Alice", 0)
		fmt.Println("Alice put result 1:", ok_1)
	}()

	go func() {
		defer wg.Done()
		ok_1 := kv.put("Bob", 0)
		fmt.Println("Bob put result 1:", ok_1)
	}()

	wg.Wait()
}
