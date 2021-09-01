package Task4

func caesar(offset int, input string) (hasil string) {
	var tmp int
	for _, val := range input {
		tmp = (int(val)-97+offset)%26 + 97
		hasil += string(tmp)
	}
	return
}
