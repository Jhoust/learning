package main

import "fmt"

func main() {

	//var a int
	//var b string
	//var c float64
	//var d bool
	var a int = 42
	var b string = "hello"
	var c float64 = 64.67
	var p bool = true
	//%v will give you the 0 value of your variables which, when printed are 0 for int, nothing for string, 0 for float, false for bool.
	//%v will give the value of the variable (when initialized) and it's a lot of code packed into two little characters.

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", p)



	//var z string = "j"
	//var x string = "a"
	//var c string = "r"
	//var v string = "e"
	//var b string = "d"

	//fmt.Println(z,x,c,v,b)

}

