package Task4

import (
	"fmt"
	"strings"
)

func compare(a, b string) (hasil string) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Print("Error: ", err)
		}
	}()

	if a < b {
		if strings.Contains(b, a) {
			hasil = a
		} else {
			panic("Bukan merupakan substr")
		}
	} else {
		if strings.Contains(a, b) {
			hasil = b
		} else {
			panic("Bukan merupakan substr")
		}
	}
	return
}
