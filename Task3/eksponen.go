package Task3

func eksponen(a int, b int) (hasil int) {
	hasil = 1
	for b > 0 {
		hasil *= a
		b--
	}
	return
}
