package main

import "fmt"

func fibonacci(end int) <-chan int {
	ch := make(chan int)
	go func() {
		a, b := 0, 1
		ch <- a
		for b < end {
			ch <- b
			a, b = b, a+b
		}
		close(ch) // Close to end the client loop
	}()
	return ch
}

func main() {
	for n := range fibonacci(1000) {
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}
