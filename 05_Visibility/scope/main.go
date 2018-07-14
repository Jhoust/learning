package main

import "fmt"

var x = "boss hogg"

func one() {
	fmt.Println(x)
}

func leftover() {
	fmt.Println(x)
}

func main() {
	leftover()
	one()
}
