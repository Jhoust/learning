package main

import "fmt"

func main() {
	if false { // a switch that will change depending on if true (first statement) or false (second statement)
		fmt.Println("this is the first statement")
	} else if true {
		fmt.Println("the second statement")
	} else {
		fmt.Println("third statemnt")

	}

}
