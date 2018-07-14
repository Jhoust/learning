package main

import "fmt"

func main() {
	greet("Jared")
	greet("Royce")
	greet("bella")

}


func greet(name string){

	fmt.Println(name)
}