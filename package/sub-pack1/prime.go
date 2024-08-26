package prime

import (
	"fmt"
)

func Prime(input int) {
	count := 0
	for a := 2; a < input/2; a++ {
		if input%a == 0 {
			count++
		}
	}
	if count == 0 {
		fmt.Println(input, "It is a Prime number")
	} else {
		fmt.Println(input, "it is Not a Prime number")
	}
}
