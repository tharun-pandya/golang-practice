package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Println("Enter a number")
	fmt.Scan(&num)
	fmt.Println(num, "is Root number of", root(num))
}
func root(data int) (root int) {
	root = data * data
	return
}
