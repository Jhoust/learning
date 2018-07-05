package main

import "fmt"

const (
	_   = iota
	KB = 1 << (iota * 10) // this is called bit shifting and it is using iota to move the value of the number
	MB = 1 << (iota * 10)//along the binary powers table, it's moving iota over 2 (binary) and adding to the power of 10 each time its called in const
	GB = 1 << (iota * 10)//iota continues to add (not multiplying but adding) 10 powers to itself
	test =  GB >> (1 * 10)+5 //this is how to move back to other decimal places
)

											// kb is 2 to the power of 10
											// mb is 2 to the power of 20
			//nothing								// gb is 2 to the power of 30

func main() {
	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)
	fmt.Printf("%b\t", test)
	fmt.Printf("%d\n", test)
}