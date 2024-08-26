package main

import "fmt"

func main() {
	var r1, r2 float32
	func() {
		fmt.Println("Enter first Ratio number")
		fmt.Scan(&r1)
		fmt.Println("Enter second Ratio number")
		fmt.Scan(&r2)
		ratio1 := (r1 + r2) / r1
		ratio2 := (r1 + r2) / r2
		fmt.Printf("ratio of First Number is %v \n", ratio1)
		fmt.Printf("ratio of Second Number is %v  ", ratio2)
	}()
}
