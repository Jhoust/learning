package main

import "fmt"

func main() {
	fmt.Println(greet("Jared ","Royce"))

}

func greet(fname,lname string)string{
	return fmt.Sprint(fname,lname)
}