package main

import "fmt"

func main() {
	var by byte
	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	var ui uint
	var ui8 uint8
	var ui16 uint16
	var ui32 uint32
	var ui64 uint64
	var f32 float32
	var f64 float64
	var co64 complex64
	var co128 complex128
	var p uintptr
	var run rune
	var bo bool
	fmt.Println("Byte ", by)
	fmt.Println("signed integer ", i)
	fmt.Println("signed integer 8bit ", i8)
	fmt.Println("signed integer 16bit ", i16)
	fmt.Println("signed integer 32bit ", i32)
	fmt.Println("signed integer 64bit ", i64)
	fmt.Println("unsigned integer ", ui)
	fmt.Println("unsigned integer 8bit ", ui8)
	fmt.Println("unsigned integer 16bit ", ui16)
	fmt.Println("unsigned integer 32bit ", ui32)
	fmt.Println("unsigned integer 64bit ", ui64)
	fmt.Println("float 32bit ", f32)
	fmt.Println("float 64bit ", f64)
	fmt.Println("complex 64bit ", co64)
	fmt.Println("complex 128bit ", co128)
	fmt.Println("pointer ", p)
	fmt.Println("rune ", run)
	fmt.Println("boolean ", bo)
}
