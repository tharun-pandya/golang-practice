package ratio

import "fmt"

func Ratio(r1, r2 float32) {
	ratio1 := (r1 + r2) / r1
	ratio2 := (r1 + r2) / r2
	fmt.Printf("ratio of First Number is %v \n", ratio1)
	fmt.Printf("ratio of Second number is %v \n", ratio2)
}
