package main

import (
	"fmt"
	"time"
)

func sum(c chan int, a, b int) {
	fmt.Println("goroutine-1 is excution started")
	sum := a + b
	c <- sum
	fmt.Println("goroutine-1 is exit")
}
func main() {
	start := time.Now()
	ch := make(chan int, 2) // unbuffer channel
	var num1, num2 int
	fmt.Println("Enter first number ")
	fmt.Scan(&num1)
	fmt.Println("Enter second number ")
	fmt.Scan(&num2)
	go sum(ch, num1, num2)
	val := <-ch
	fmt.Println(val)
	go mul(ch, num1, num2)
	val = <-ch
	fmt.Println(val)
	fmt.Println(time.Since(start))
}
func mul(c chan int, a, b int) {
	fmt.Println("goroutine-2 is excution started")
	mul := a * b
	c <- mul
	fmt.Println("goroutine-2 is exit")
}
