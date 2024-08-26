package main

import (
	"fmt"
)

func main() {
	// defer follows First in Last Out (stack)
	defer fmt.Println("four")  //last
	defer fmt.Println("three") // third
	defer fmt.Println("two")   //second
	defer fmt.Println("one")   // first
	table()
}
func table() {
	defer fmt.Println("Table close") // it will be execute end of the program
	var num int
	fmt.Println("Enter a Number for to Print table ")
	fmt.Scan(&num)
	for i := 1; i <= 10; i++ {
		fmt.Println(num, "x", i, "=", num*i)
	}
}
