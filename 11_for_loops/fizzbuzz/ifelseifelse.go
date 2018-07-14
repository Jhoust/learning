package main   //works but prints twice on 5 and 10 as fizz and buzz with fizzbuzz

import "fmt"

func main() {
	i := 1
	for {
		if i%2 == 1 {
			fmt.Println("fizz -", i)
		}
		if i%2 == 0 {
			fmt.Println("buzz -", i)
		}
		if i%5 == 0 {
			fmt.Println("fizzbuzz", i)

		} else if i > 10 {
			break
		}
		i++
	}
}
