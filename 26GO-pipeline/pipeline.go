package main

import "fmt"

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range nums {
			out <- v
		}
		close(out)
	}()

	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	// // Set up the pipeline.
	// c := gen(2, 3, 5)
	// out := sq(c)

	// // Consume the output.
	// fmt.Println(<-out) // 4
	// fmt.Println(<-out) // 9

	// Set up the pipeline and consume the output.
	for n := range sq(sq(gen(2, 3, 1, 34, 45, 56))) {
		fmt.Println(n) // 16 then 81
	}
}
