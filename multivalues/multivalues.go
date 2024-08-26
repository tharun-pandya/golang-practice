package main

import (
	"fmt"
)

func main() {
	var a, b, c, d int = 34, 56, 2, 45 //multi varibules
	const (
		A1 = 10
		B1 = 3
	) // multi constants
	fmt.Println("Sum of ", a, " and ", b, " is =", a+b)
	fmt.Println("Product of ", c, " and ", d, " is =", c*d)
	fmt.Println("mod of ", A1, " and ", B1, " is =", A1%B1)
	varibules()
	constants()
}
func varibules() {
	var a, b, option int
	fmt.Println("Enter first Number ")
	fmt.Scan(&a)
	fmt.Println("Enter Second Number ")
	fmt.Scan(&b)
	fmt.Println("Sum of ", a, " and ", b, " is =", a+b)
	fmt.Println("Product of ", a, " and ", b, " is =", a*b)
	fmt.Println("mod of ", a, " and ", b, " is =", a%b)
	fmt.Println("do you want to change data data \n yes for enter 1 or no for enter 0")
	fmt.Scan(&option)
	for option == 1 {
		fmt.Println("Update first Number ")
		fmt.Scan(&a)
		fmt.Println("Update Second Number ")
		fmt.Scan(&b)
		fmt.Println("Sum of ", a, " and ", b, " is =", a+b)
		fmt.Println("Product of ", a, " and ", b, " is =", a*b)
		fmt.Println("mod of ", a, " and ", b, " is =", a%b)
		fmt.Println("do you want to change data data \n yes for enter 1 or no for enter 0")
		fmt.Scan(&option)
	}
}
func constants() {
	const (
		A = 50
		B = 6
	)
	var a, b, option int
	fmt.Println("Sum of ", A, " and ", B, " is =", A+B)
	fmt.Println("Product of ", A, " and ", B, " is =", A*B)
	fmt.Println("mod of ", A, " and ", B, " is =", A%B)
	fmt.Println("do you want to change data data \n yes for enter 1 or no for enter 0")
	fmt.Scan(&option)
	for option == 1 {
		fmt.Println("Update first Number ")
		_, err1 := fmt.Scanln(&a)
		fmt.Println("Update Second Number ")
		_, err2 := fmt.Scanln(&b)
		if err1 == nil {
			fmt.Println(err1, "Constants are not chnaged at run time")
		}
		if err2 == nil {
			fmt.Println(err2, "Constants are not chnaged at run time")
			break
		}
	}
	return

}
