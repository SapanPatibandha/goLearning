package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Welcome to chanels in golang.")

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// myCh <- 5
	// fmt.Println(<-myCh)

	wg.Add(2)
	go func(ch chan int, w *sync.WaitGroup) {
		val, isChanelOpen := <-myCh

		if isChanelOpen {
			fmt.Println(val)
		} else {
			fmt.Println("Chanel close")
		}
		// fmt.Println(<-myCh)

		wg.Done()
	}(myCh, wg)

	go func(ch chan int, w *sync.WaitGroup) {
		// myCh <- 5
		myCh <- 0

		close(myCh)
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
