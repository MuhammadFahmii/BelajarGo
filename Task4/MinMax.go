package Task4

/*
	Sample Test Case
	a, b, c, d, e := 7, 5, 9, 2, -1
	Input: MinMax(&a, &b, &c, &d, &e)
	Output: -1 9
*/
func MinMax(number ...*int) (min, max int) {
	min = *number[0]
	for _, data := range number {
		if *data < min {
			min = *data
		} else if *data > max {
			max = *data
		}
	}
	return
}
