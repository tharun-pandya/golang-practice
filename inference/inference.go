package main

import "fmt"

func main() {
	var num1 = 98 // When declaring a variable without specifying an explicit type,  the variable's type is inferred from the value on the right hand side.
	num2 := 76    //without declaring var
	fmt.Println("sum of", num1, "and", num2, "is", num1+num2)
}
