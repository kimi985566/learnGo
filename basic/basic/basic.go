package main

import (
"fmt"
"math/cmplx"
"math"
)

//包内部变量
var (
	aa = 3
	kk = "true"
)

func variableZeroValue() {
	var a int
	var s string
	//fmt.Println(a, s)
	fmt.Printf("%d %q\n", a, s)
}

func variableInitValue() {
	var a, b int = 2, 3
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, d = 1, "abc", true, "def"
	fmt.Println(a, b, c, d)
}

func variableShorter() {
	a, b, c, d := 1, "abc", true, "def"
	fmt.Println(a, b, c, d)
}

func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func consts() {
	const (
		a, b = 3, 4
		s    = "abc.txt"
	)
	fmt.Println(a, b, s)
}

func enums() {
	const (
		b  = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, gb)
}

func main() {
	fmt.Println("Hello,world")
	variableZeroValue()
	variableInitValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, kk)

	euler()
	triangle()
	consts()
	enums()
}

