package main

import "fmt"

func main() {
	var int_fac interface{}
	var name string
	fmt.Println("Enter your name")
	fmt.Scan(&name)
	int_fac = name
	var change = int_fac.(string)
	fmt.Println("interface value is", change)
	value, test := int_fac.(int)
	if test {
		fmt.Println("integer found", value)
	} else {
		fmt.Println("integer not found")
	}
}
