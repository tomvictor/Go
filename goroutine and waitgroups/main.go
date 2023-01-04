package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		fmt.Println("Async Execution")
		wg.Done()
	}()
	wg.Wait()

	fmt.Println("Sync Execution")
}
