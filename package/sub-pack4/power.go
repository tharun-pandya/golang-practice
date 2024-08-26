package Power

import (
	"fmt"
)

func Power(in1, in2 int) {
	var i, in int
	in = 1
	for i = 0; i < in2; i++ {
		in = in * in1
	}
	fmt.Println(in1, "Power", in2, "is", in)
}
