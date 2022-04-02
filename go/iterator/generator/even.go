package main

import "fmt"

func integers(end int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < end; i++ {
			ch <- i
		}
		close(ch) // Close to end the client loop
	}()
	return ch
}

func evens(generator <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := range generator {
			if i%2 == 0 {
				ch <- i
			}
		}
		close(ch) // Close to end the client loop
	}()
	return ch
}

func main() {
	for n := range evens(integers(10)) {
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}
