package main

import (
	"strconv"
)

func main() {
	var i int = 76
	var f float32 = 98.76
	var s string = "123456"
	var b bool = true
	var s1 string = "false"
	var t1 float32 = float32(i)
	var t2, _ int = strconv.ParseInt(s, 10, 64)
}
