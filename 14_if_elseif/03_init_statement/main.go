package main

import "fmt"

func main() {

	b := true

	if food := "choclate"; b { //initialization statement being declared in this tight program scope
		fmt.Println(food) // this says two things food is given the string value of "chocolate" and; (b) the previous statement was true, go ahead and execute the next command
	}
	fmt.Println(food) //the variable food is out of scope
}

//package main
//
//import "fmt"
//
//func main() {
//
//	b := true
//
//	if  b{
//		fmt.Println("food") // the string will execute on its own
//	}
//}
