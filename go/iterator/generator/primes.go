package main

import "fmt"

type primeCell struct {
	N, M int
}

func primes(end int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch) // Close to end the client loop
		if end <= 2 {
			return
		}
		ch <- 2
		if end <= 3 {
			return
		}
		ch <- 3
		if end <= 5 {
			return
		}
		primes := []primeCell{{3, 3}}
		for i := 5; i < end; i += 2 {
			isprime := true
			for _, p := range primes {
				for p.M < i {
					p.M += p.N
				}
				if p.M == i {
					isprime = false
					break
				}
			}
			if isprime {
				primes = append(primes, primeCell{i, i})
				ch <- i
			}
		}

	}()
	return ch
}

func main() {
	for end := 0; end < 21; end++ {
		fmt.Printf("end = %3d: ", end)
		for n := range primes(end) {
			fmt.Printf("%d ", n)
		}
		fmt.Println()
	}
}
