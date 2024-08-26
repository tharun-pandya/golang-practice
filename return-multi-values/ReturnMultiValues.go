package main

import "fmt"

func main() {
	var data int
	fmt.Println("Enter a number to find \n 1.Cube volume \n 2.Lateral Surface Area \n 3.Total Surface Area")
	fmt.Scanln(&data)
	a, b, c := cube(data)
	fmt.Println("Cube Volume is", a)
	fmt.Println("Cube Lateral Surface Area(LSA) is", b)
	fmt.Println("Cube Total Surface Area(TSA) is", c)

}
func cube(num int) (volume, lsa, tsa int) {
	volume = num * num * num
	lsa = 4 * num * num
	tsa = 6 * num * num
	return
}
