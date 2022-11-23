package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	/*
		// solution 1
		var mu sync.Mutex
		count := 0
	*/

	//count := int64(0)
	var count int64
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {

			defer wg.Done()
			for j := 0; j < 10_000; j++ {
				/* solution 1
				mu.Lock()
				count++
				mu.Unlock()*/
				atomic.AddInt64(&count, 1)
			}
		}()
	}

	wg.Wait()

	fmt.Println(count)
}
