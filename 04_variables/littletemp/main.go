package main

import "fmt"

func main() {
	var input float64
	var f float64
	fmt.Println("Give me you temperature in celsius")
	fmt.Scan(&input)
	f = input*1.8 + 32
	fmt.Println(f, "Fahrenheit")
}
