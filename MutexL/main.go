package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.RWMutex

//
// func inc(x *int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	mutex.Lock()
// 	*x++
// 	mutex.Unlock()
// }
//
// func main() {
// 	x := 0
// 	wg := sync.WaitGroup{}
// 	for i := 0; i < 1000; i++ {
// 		wg.Add(1)
// 		go inc(&x, &wg)
// 	}
// 	wg.Wait()
// 	log.Printf("final value: %d\n", x)
// }

// a simple function that returns true if a number is even
func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	n := 0

	// goroutine 1
	// reads the value of n and prints true if its even
	// and false otherwise
	go func() {
		mutex.RLock()
		defer mutex.RUnlock()

		nIsEven := isEven(n)
		// we can simulate some long running step by sleeping
		// in practice, this can be some file IO operation
		// or a TCP network call
		time.Sleep(5 * time.Millisecond)
		if nIsEven {
			fmt.Println(n, " is even")
			return
		}
		fmt.Println(n, "is odd")
	}()

	go func() {
		mutex.RLock()
		defer mutex.RUnlock()

		nIsPositive := n > 0
		if nIsPositive {
			fmt.Println(n, "is positive")
			return
		}
		fmt.Println(n, "is negative")
	}()

	// goroutine 2
	// modifies the value of n
	go func() {
		mutex.Lock()
		defer mutex.Unlock()
		n++
	}()

	// just waiting for the goroutines to finish before exiting
	time.Sleep(time.Millisecond * 100)
}
