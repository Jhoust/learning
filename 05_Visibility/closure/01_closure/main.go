package main

import "fmt"

var x = 0

func increment() int {
	//x = x + 1
	//x += 1         //all 3 of these do the same thing, returning x and then calling the increment function to print it out stops it from looping
	x++
	return x
}
	func main() {

		fmt.Println(increment())   //when calling the (increment()) function be sure to include () otherwise you run into strange errors
		fmt.Println(increment())   //need both print lines to show the addition of the program (x++ is) = (x = x+1) not x+x

	}