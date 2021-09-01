package Task2

import "fmt"

func playWithAsterik(bil int) {
	for i := 1; i <= bil; i++ {
		for j := bil; j > i; j-- {
			fmt.Print(" ")
		}
		for k := 1; k <= i; k++ {
			fmt.Print("*")
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
