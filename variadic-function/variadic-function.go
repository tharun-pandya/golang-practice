package main

import (
	"fmt"
)

func main() {
	var num, in1, in2 int
	var add1 []int
	var add2 []int
	fmt.Println("Enter how many pair number you want to add")
	fmt.Scan(&num)
	for a := 0; a < num; a++ {
		fmt.Println("Enter First number in", a+1, "pair")
		fmt.Scan(&in1)
		add1 = append(add1, in1)
		fmt.Println("Enter Second number in", a+1, "pair")
		fmt.Scan(&in2)
		add2 = append(add2, in2)
	}
	for a := 0; a < num; a++ {
		fmt.Println("Sum of ", a+1, "pair", add1[a], "and", add2[a], "is", add(add1[a], add2[a]))
	}
}
func add(num ...int) int {
	var ans int
	ans = num[0] + num[1]
	return ans
}
