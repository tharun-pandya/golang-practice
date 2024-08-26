package percentage

import "fmt"

func Percentage(cp, sp int) {
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
