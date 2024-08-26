package main

import "fmt"

func main() {
	var a [5]int
	var b = [7]int{3, 4, 12, 76, 1}      // automatically allocates memory to 7 elements
	var c = [...]int{76, 98, 45}         //only memory allocates to these three elements
	var d = [9]int{2: 76, 3: 98, 7: 125} //data stored in particular position
	fmt.Println("enter only 5 elements in first array")
	for i := 0; i < len(a); i++ {
		fmt.Println("Enter an element in the array")
		fmt.Scan(&a[i])
	}
	for i := 0; i < len(a); i++ {
		fmt.Println("first array is", a[i])
		fmt.Printf("\n")
	}
	for i := 0; i < len(b); i++ {
		fmt.Println("second array is", b[i])
	}
	fmt.Printf("\n")
	for i := 0; i < len(c); i++ {
		fmt.Println("third array is", c[i])
	}
	fmt.Printf("\n")
	for i := 0; i < len(d); i++ {
		fmt.Println("last array is", d[i])
	}
}
