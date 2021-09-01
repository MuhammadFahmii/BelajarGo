package Task6

import (
	"fmt"
	"math"
)

func SimpleEquation(a, b, c int) {
	n := c
	x, y, z := 0, 0, 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			for k := 1; k <= n; k++ {
				if i+j+k == a && i*j*k == b && math.Pow(float64(i), 2)+math.Pow(float64(j), 2)+math.Pow(float64(k), 2) == float64(c) {
					x, y, z = k, j, i
				}
			}
		}
	}
	if x == 0 && y == 0 && z == 0 {
		fmt.Println("No Solutions")
	} else {
		fmt.Println(x, y, z)
	}
}
