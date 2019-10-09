package intmath

// Gcd returns the greatest common denominator of a, b
func Gcd(a, b int) int {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}
	return a
}
