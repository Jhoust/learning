package main

import "fmt"

func main() {
	i := 0
	for {
		i++

		if i%2 == 0 {
			fmt.Println("buzz")
		}
		if i%2 == 1 {
			fmt.Println("fizz")
		}
		if i%5 == 0 {
			fmt.Println("fizzbuzz")

			continue
		}

		if i > 50 {
			break
		}

	}
}