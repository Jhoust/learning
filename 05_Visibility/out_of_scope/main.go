package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	{
		y := "yo what's up"
		//var y = "Hello"
		fmt.Println(y)
	}
	//fmt.Println(y) out of scope
}
