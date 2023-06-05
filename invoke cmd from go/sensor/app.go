package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for i := 1; i < 100; i++ {
			fmt.Printf("value: %v\n", i)
			time.Sleep(time.Second * 1)
		}
	}()
	wg.Wait()
}
