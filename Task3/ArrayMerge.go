package Task3

func duplikat(arr []string, a string) bool {
	for _, value := range arr {
		if value == a {
			return true
		}
	}
	return false
}

func Gabung(arrayA, arrayB []string) (hasil []string) {
	var x []string = append(arrayA, arrayB...)
	for _, value := range x {
		if !duplikat(hasil, value) {
			hasil = append(hasil, value)
		}
	}
	return
}
