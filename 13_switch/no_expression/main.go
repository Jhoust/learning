package main

import "fmt"

func main () {
	 Mywife := "amber"

	switch  {								//no expression is listed here
	case (Mywife) == "Jared":				//parenthesis needed
		fmt.Println("Hello,Jared")
	case (Mywife) == "amber":
		fmt.Println("Hello my love")
	case (Mywife) == "royce":
		fmt.Println("goodboy!")
		fallthrough							//fallthrough is unused because answer was found before
	case (Mywife) == "bella":
		fmt.Println("goodgirl!")
	default:								//default a good idea
		fmt.Println("I am default")
	}
}