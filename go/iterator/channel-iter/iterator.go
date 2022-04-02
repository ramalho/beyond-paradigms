package main

func PositiveIterator(s []float64) <-chan float64 {
	ch := make(chan float64)
	go func() {
		for _, v := range s {
			if v > 0 {
				ch <- v
			}
		}
		close(ch) // terminate client loop
	}()
	return ch
}
