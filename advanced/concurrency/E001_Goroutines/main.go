package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func fiveMilliDelay() {
	for i := 0; i < 10; i++ {
		fmt.Println("fiveMilliDelay,  ->", i)
		time.Sleep(time.Duration(5 * time.Millisecond))
	}
	wg.Done()
}

func tenMilliDelay() {
	for i := 0; i < 10; i++ {
		fmt.Println("tenMilliDelay,  ->", i)
		time.Sleep(time.Duration(10 * time.Millisecond))
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go fiveMilliDelay()
	go tenMilliDelay()
	wg.Wait()
}
