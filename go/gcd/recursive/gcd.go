package intmath

// Gcd returns the greatest common denominator of a, b
func Gcd(a, b int) int {
	if a == b {
		return a
	}
	if a > b {
		return Gcd(b, a-b)
	}
	return Gcd(a, b-a)
}
