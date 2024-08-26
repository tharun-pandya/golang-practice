package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter number to Find Root")
	fmt.Scan(&num)
	answer := squreroot(num)
	if answer == 0 {
		fmt.Println("Invalid Number", num, " This number doesn't have any Root")
	} else {
		fmt.Println("Root Number of", num, "is", answer)
	}
}

// single parameter and single return value
func squreroot(root int) int {
	for a := 0; a <= root/2; a++ {
		result := a * a
		if result > root {
			return 0
		} else if result == root {
			return a
		}
	}
	return 0
}
