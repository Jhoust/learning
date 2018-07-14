package main

import "fmt"

func main() {
	fmt.Println(greet("Jared ","Royce "))

}

func greet(fname string, lname string)(string ,string){
	return fmt.Sprint(lname,fname), fmt.Sprint(fname,lname)
}