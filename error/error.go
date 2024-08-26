package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var i, in int64
	a := bufio.NewReader(os.Stdin)
	b := bufio.NewReader(os.Stdin)
	fmt.Println("Enter First number")
	input1, err := a.ReadString('\n')
	in1, err := strconv.ParseInt(strings.TrimSpace(input1), 10, 64)
	fmt.Println("Enter Second Number")
	input2, err := b.ReadString('\n')
	in2, err := strconv.ParseInt(strings.TrimSpace(input2), 10, 64)
	if err != nil {
		panic(err)
	} else {
		in = 1
		for i = 0; i < in2; i++ {
			in = in * in1
		}
		fmt.Println(in1, "Power", in2, "is", in)
	}
}
