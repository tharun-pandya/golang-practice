package main

import "fmt"

func main() {
	var cp, sp int
	fmt.Println("Enter Cost of Product you Purchase")
	fmt.Scan(&cp)
	fmt.Println("Enter You Selling Price")
	fmt.Scan(&sp)
	a, b, c := func(c, s int) (val, per int, str string) {
		if c > s {
			str = "Loss"
			val = c - s
			per = ((c - s) * 100) / c
			return
		} else if c < s {
			str = "Profit"
			val = s - c
			per = ((s - c) * 100) / c
			return
		}
		return
	}(cp, sp)
	fmt.Printf("%v in Price is %vRs\n", c, a)
	fmt.Printf("%v in Percentage is %v%%", c, b)
}
