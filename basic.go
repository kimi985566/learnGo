package main

import (
	"fmt"
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

func main() {
	fmt.Println("Hello,world")
	variableZeroValue()
	variableInitValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(aa, kk)
}
