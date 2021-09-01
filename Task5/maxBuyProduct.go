package Task5

import "fmt"

/*
	Sample Test Cases
	Input: money = 50000, productPrice = [25000, 25000, 10000, 14000]
	Output: 3

	Input: Input: money = 30000, productPrice = [15000, 10000, 12000, 5000, 3000]
	Output: 4
*/
func MaxBuyProduct(money int, productPrice []int) {
	n, jumlahBarang, jumlahBayar := len(productPrice), 0, 0
	for k := range productPrice {
		min := k
		for j := k + 1; j < n; j++ {
			if productPrice[j] < productPrice[min] {
				min = j
			}
		}
		productPrice[k], productPrice[min] = productPrice[min], productPrice[k]
	}

	for i := range productPrice {
		if jumlahBayar+productPrice[i] > money {
			fmt.Println(jumlahBarang)
			break
		} else if jumlahBayar+productPrice[i] == money {
			jumlahBarang++
			fmt.Println(jumlahBarang)
			break
		} else {
			jumlahBayar += productPrice[i]
			jumlahBarang++
		}
	}
}
