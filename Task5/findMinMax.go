package Task5

import "fmt"

/**
Sample test case
Input= [5, 7, 4, -2, -1, 8]
Output= min: -2 index: 3 max: 8 index: 5

Input= [2, -5, -4, 22, 7, 7]
Output= min: -5 index: 3 max: 22 index: 3
*/
func FindMinMax(arr []int) string {
	indexMin, indexMax, min, max := 0, 0, arr[0], 0
	for i, data := range arr {
		if data < min {
			indexMin = i
			min = data
		} else if data > max {
			indexMax = i
			max = data
		}
	}
	return fmt.Sprintf("min: %d index: %d max: %d index: %d", min, indexMin, max, indexMax)
}
