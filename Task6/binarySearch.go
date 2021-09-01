package Task6

/*
	fmt.Println(Task6.BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 53))
*/
func BinarySearch(arr []int, x int) int {
	l, n := 0, len(arr)-1
	for l <= n {
		m := l + (n-l)/2.0
		if arr[m] == x {
			return m
		} else if arr[m] < x {
			l = m + 1
		} else {
			n = m - 1
		}
	}
	return -1
}
