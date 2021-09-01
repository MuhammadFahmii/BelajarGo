package main

import (
	"BelajarGo/Task6"
	"fmt"
)

func main() {
	// fmt.Println(Task6.MoneyCoins(123))
	// fmt.Println(Task6.MoneyCoins(543))
	fmt.Printf("Hasil: %d\n", Task6.BinarySearch([]int{1, 1, 3, 5, 5, 6, 7}, 3))
	fmt.Printf("Hasil: %d", Task6.BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 53))
}
