package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	var str1 string = "Hello"
	var str2 = "World"
	fmt.Println("start")
	for i := 0; i < 10; i++ {
		go hello(str1)
		hello(str2)
	}
	go func(s string) {
		time.Sleep(1 * time.Second)
		fmt.Println(s)
	}("going")
	time.Sleep(5 * time.Second)
	fmt.Println("program exit")
	fmt.Println(time.Since(start))

}
func hello(str string) {
	time.Sleep(3 * time.Second)
	fmt.Println(str)
	if str == "World" {
		fmt.Println("")
	}
}
