package Task2

func pangkat(base, pangkat int) int {
	hasil := 1
	for pangkat > 0 {
		hasil *= base
		pangkat--
	}
	return hasil
}
