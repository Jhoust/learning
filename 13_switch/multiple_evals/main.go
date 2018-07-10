package main

import "fmt"

func main () {
	switch "Eric"{

	case "bella", "royce":						//multiple evals are okay
		fmt.Println("good puppers!")

	case "Jared" , "Amber":
		fmt.Println("Hard at work")

	case "Eric", "screaming little girl next door":
	fmt.Println("you guys need to be quite")

	}
}
