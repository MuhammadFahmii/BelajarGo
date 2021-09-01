package Task5

/*
	Dalam matematika, bil prima adalah bil asli yang lebih besar dari angka 1, yang faktor pembaginya adalah 1 dan bil itu sendiri. Buat fungsi yang menampilkan bil prima sesuai dengan deret urutannya

	Sample Test Case
	Input: 1
	Output: 2

	Input: 5
	Output: 11
*/
func PrimeX(number int) int {
	bil := 2
	for {
		faktor := 0
		for i := 1; i <= bil; i++ {
			if bil%i == 0 {
				faktor++
			}
		}
		if faktor == 2 {
			number--
		}
		if number == 0 {
			return bil
		}
		bil++
	}
}
