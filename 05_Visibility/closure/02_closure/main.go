package main

import "fmt"

var x = 0

func main() {

	increment := func() int { //the function needs to be set so (int) is needed after declaring the name
		x++
		return x // this is the way to add a function inside another function. The name of the function needs to be
		// this is called a "func expression"; when you assign a func to a variable otherwise it would be an anon func
	}

	fmt.Println(increment())
	fmt.Println(increment())
}
