package Task2

import "fmt"

func cetakTablePerkalian(bil int) {
	for i := 1; i <= bil; i++ {
		for j := 1; j <= bil; j++ {
			fmt.Print(j * i)
			fmt.Print(" ")
		}
		fmt.Println(" ")
	}
}
