package main   //works but not the way its suppose too 5 and 10

import "fmt"

func main() {
	i := 0
	for {
		i++ // adds 1 to zero immediately

		if i%2 == 0 {
			fmt.Println("buzz") //this code almost works, but creates extra fizz(5) or buzz(10) to fizzbuzz and tehn an extra fizz at the end
		}
		if i%2 == 1 {
			fmt.Println("fizz")
		}
		if i%5 == 0 {
			fmt.Println("fizzbuzz")
			continue
		}

		if i >= 10 {
			break
		}
	}

}
