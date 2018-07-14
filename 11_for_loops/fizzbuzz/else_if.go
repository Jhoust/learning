package main   //works

import "fmt"

func main() {
											//this code works and the i has to be set to 1 otherwise it will print "fives first"
	for i := 1; i <= 10;i++  {
		if i%5 == 0 {
			fmt.Println("fives -", i)
		} else if i%2 == 1 {
			fmt.Println("odd -", i)
		} else if i%2 == 0 {
			fmt.Println("even -", i)
		}else {
			fmt.Println(i)
		}

	}
}
