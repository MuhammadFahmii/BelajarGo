package Task5

import "fmt"

func MostAppear(items []string) (hasil []string) {
	totalPerItem := map[string]int{}
	for _, val := range items {
		totalPerItem[val] += 1
	}
	min := totalPerItem[items[0]]
	fmt.Println(min)
	for i, val := range totalPerItem {
		if val < min {
			hasil = append(hasil, i)
		}
	}
	fmt.Println(totalPerItem)
	return
}
