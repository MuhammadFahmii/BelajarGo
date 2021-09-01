package Task4

func Swap(a, b *int) {
	*a, *b = *b, *a

	// Cara lama
	// tmp := *b
	// *b = *a
	// *a = tmp
}
