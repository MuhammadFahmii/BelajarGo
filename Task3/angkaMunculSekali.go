package Task3

import (
	"strconv"
)

func munculSekali(str string) (hasil []int) {
	jumlahMirip := map[int]int{}
	for _, val := range str {
		konversi, _ := strconv.Atoi(string(val))
		jumlahMirip[konversi] += 1
	}
	for key, val := range jumlahMirip {
		if val == 1 {
			hasil = append(hasil, key)
		}
	}
	return
}
