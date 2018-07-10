package main

import "fmt"



func main() {

    var x uint64 = 0
	var action uint8 = 0


	fmt.Println("Enter a whole number:")
	fmt.Scan(&action)

    x = 1 << (action + 1)

	fmt.Println(x)
}