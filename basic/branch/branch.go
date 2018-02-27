package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	const file = "abc.txt"
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}
