package Task2

func palindrome(kata string) bool {
	var palindrome = ""
	for i := len(kata) - 1; i >= 0; i-- {
		palindrome += string(kata[i])
	}
	if kata == palindrome {
		return true
	} else {
		return false
	}
}
