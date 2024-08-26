package main

import (
	"fmt"
)

func main() {
	prime()
}
func prime() {
	var input int
	fmt.Println("enter a number checking Prime or Not")
	fmt.Scanf("%v", &input)
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
