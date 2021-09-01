package Task2

import "math"

func PrimeNumber(bil int) bool {
	if bil <= 1 {
		return false
	}
	for i := 1; float64(i) <= math.Sqrt(float64(bil)); i++ {
		if i == 1 {
			i++
		}
		if bil%i == 0 {
			return false
		}
	}
	return true
}
