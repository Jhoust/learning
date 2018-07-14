package main

import "fmt"

func main() {
	switch "amber" {

	case "Jared":
		fmt.Println("Hello,Jared")

	case "amber":
		fmt.Println("Hello my love")

	case "royce":
		fmt.Println("goodboy!")

		fallthrough

	case "bella":
		fmt.Println("goodgirl!")
	}
}
