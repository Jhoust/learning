package main

import "fmt"

func main() {
	fmt.Println(greet("Jared ","Royce"))


}


func greet(fname string,lname string)(s string) {
	s = fmt.Sprint(fname,lname)
	return

}