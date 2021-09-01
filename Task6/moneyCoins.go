package Task6

func MoneyCoins(money int) (hasil []int) {
	nominal := []int{1, 10, 20, 50, 100, 200, 500, 1000, 2000, 5000, 10000}
	for i := len(nominal) - 1; i >= 0; i-- {
		for money >= nominal[i] {
			money -= nominal[i]
			hasil = append(hasil, nominal[i])
		}
	}
	return
}
