package Task5

/**
Sample Test Cases
Input [-2, 1, -3, 4, -1, 2, 1, -5, 4]
Output : 6
Penjelasan : 6 adalah hasil penambahan dari deret 4, -1, 2, 1

Input: [-2, -5, 6, -2, -3, 1, 5, -6]
Output: 7
Penjelasan: 7 adalah hasil penambahan dari deret 6, -2, -3, 1, 5
*/

func MaxSequence(arr []int) (hasil int) {
	hasil, min := 0, arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] <= min {
			min = arr[i]
			if i == len(arr)-1 || i == len(arr)-2 {
				if arr[len(arr)-1] > hasil {
					return arr[len(arr)-1]
				}
				return
			}
			hasil = 0
		} else {
			hasil += arr[i]
		}
	}
	return
}
