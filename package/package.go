package main

import (
	"fmt"
	p "package/sub-pack1"
	r "package/sub-pack2"
	per "package/sub-pack3"
	pow "package/sub-pack4"
	ro "package/sub-pack5"
)

func main() {
	for {
		var i, num1, num2 int
		fmt.Println("0.exit")
		fmt.Println("1.Prime")
		fmt.Println("2.ratio")
		fmt.Println("3.Percentage")
		fmt.Println("4.Add")
		fmt.Println("5.Power")
		fmt.Println("Choose One Option")
		fmt.Scan(&i)
		if i == 0 {
			break
		} else if i == 1 {
			fmt.Println("enter a number checking Prime or Not")
			fmt.Scan(&num1)
			p.Prime(num1)
		} else if i == 2 {
			var r1, r2 float32
			fmt.Println("Enter first Ratio number")
			fmt.Scan(&r1)
			fmt.Println("Enter second Ratio number")
			fmt.Scan(&r2)
			r.Ratio(r1, r2)
		} else if i == 3 {
			fmt.Println("Enter Cost of Product you Purchase")
			fmt.Scan(&num1)
			fmt.Println("Enter You Selling Price")
			fmt.Scan(&num2)
			per.Percentage(num1, num2)
		} else if i == 4 {
			fmt.Println("Enter First number")
			fmt.Scan(&num1)
			fmt.Println("Enter Second Number")
			fmt.Scan(&num2)
			pow.Power(num1, num2)
		} else if i == 5 {

			fmt.Println("Enter a number")
			fmt.Scan(&num1)
			ro.Root(num1)
		} else {
			fmt.Println("Invalid choice please enter Index value")
		}
	}
}
