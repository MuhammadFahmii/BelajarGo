package Task3

func PairSum(arr []int, target int) (int, int) {
	if len(arr) <= 1 {
		return -1, -1
	}
	m := make(map[int]int, len(arr))
	for i, v := range arr {
		j, ketemu := m[v]
		if ketemu {
			return j, i
		}
		m[target-v] = i
	}
	return -1, -1
}
