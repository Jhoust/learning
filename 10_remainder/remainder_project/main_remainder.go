package main

import "fmt"


func main(){
var x int
var y int
var z int
var w int

	fmt.Println("Give me two numbers to divide")
	fmt.Println("Type the dividend now:")
	fmt.Scanln(&x)
	fmt.Println("Type the divisor now:")
	fmt.Scanln(&y)
z = x / y
fmt.Println("The quotient is:")
	fmt.Println(z)
fmt.Println("With a remainder of: ")
w = x % y
fmt.Println(w)


	}
