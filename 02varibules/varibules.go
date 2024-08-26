package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	a := bufio.NewReader(os.Stdin)
	b := bufio.NewReader(os.Stdin)
	fmt.Println("Enter first number")
	input1, _ := a.ReadString('\n')
	a1, _ := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	fmt.Println("Enter second number")
	input2, _ := b.ReadString('\n')
	b1, _ := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	fmt.Println("sum of two numbers is", a1+b1)
}
