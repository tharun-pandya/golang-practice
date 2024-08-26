package main

import (
	"fmt"
)

func isPalindrome(c chan string, x int) {
	var num, lastdig int
	temp := x
	if temp < 0 {
		c <- "it is not a palindrome"
	} else {
		for temp > 0 {
			lastdig = temp % 10
			num = num*10 + lastdig
			temp /= 10
		}
		if num == x {
			c <- "it is a palindrome"
		}
		c <- "it is not a palindrome"
	}
}
func evenOrodd(channel chan string, num int) {
	switch num % 2 {
	case 0:
		channel <- "it is a Even Number"
	case 1:
		channel <- "it is Odd Number"
	}
}
func main() {
	var x int
	var ch = make(chan string)
	var ch1 = make(chan string)
	fmt.Println("enter a number")
	fmt.Scan(&x)
	go isPalindrome(ch, x)
	go evenOrodd(ch1, x)
	select {
	case op := <-ch:
		fmt.Println(op)
	case op := <-ch1:
		fmt.Println(op)
	}
}
