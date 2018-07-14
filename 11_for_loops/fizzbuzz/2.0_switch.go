package main

import "fmt"

var a1 string
var x int

func main() {
	fmt.Println("Pick a small number: ")
	fmt.Scan(&x)

	for i := 1 ; i <= 30; i++{

		if i%x == 0 {
			a1 = "c"
		}//else if i%x == 1 {
			//a1 = "b"
			//}


		switch a1 {
		case "a":
			fmt.Println("buzz -", i)
		case "b":
			fmt.Println()//x,"does not divide into",i,"evenly")
		case "c":
			fmt.Println(x,"divides into",i,"evenly")
		default:
			fmt.Println()//"you broke it")
		}
	}
}

