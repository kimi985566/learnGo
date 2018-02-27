package main

import (
	"io/ioutil"
	"fmt"
)

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong Score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	const file = "abc.txt"
	if contents, err := ioutil.ReadFile(file); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	fmt.Println(
		grade(59),
		grade(88),
		grade(99),
		grade(100),
		grade(0),
	)
}
