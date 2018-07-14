package main

import "fmt"

const( _ = iota
a  = iota * 10 // 1 * 10  every time iota is called it adds one to the value to iota
b  = iota * 10// 2 * 10 iota is called a 3rd time and holds the value of 2
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
}