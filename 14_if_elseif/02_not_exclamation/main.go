package main

import "fmt"

func main() {

	if !true {
		fmt.Println("this didnt run")
	}
	if !false {
		fmt.Println("this ran") // the text shows you what is executing
	}
}
