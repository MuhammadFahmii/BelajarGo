package Task5

import "fmt"

/*
PrimaSegiEmpat(2, 3, 13)
17 19
23 29
31 37
156

PrimaSegiEmpat(5, 2, 1)
2 	3 	5 	7 	11
13 	17 	19 	23 	29
*/

func PrimaSegiEmpat(high, wide, start int) {
	batasAwal, batasAkhir := 0, high*wide
	primeAfterStart := []int{}
	jumlah := 0
	start += 1
	for {
		faktor := 0
		for i := 1; i <= start; i++ {
			if start%i == 0 {
				faktor++
			}
		}
		if faktor == 2 {
			primeAfterStart = append(primeAfterStart, start)
			jumlah = jumlah + start
			batasAwal++
		}
		if batasAwal == batasAkhir {
			for i := 0; i < batasAkhir; i++ {
				fmt.Printf("%d\t", primeAfterStart[i])
				if (i+1)%(batasAkhir/wide) == 0 {
					fmt.Println("")
				}
				// Cara lebih susah
				// for j := (i*high - batasPerBaris) - 1; j < batasAkhir/wide*i; j++ {
				// 	fmt.Printf("%d\t", primeAfterStart[j])
				// }
			}
			fmt.Println(jumlah)
			return
		}
		start++
	}
}
