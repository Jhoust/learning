package main  //works prints of 5 and 10 twice as buzz and fizz

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		if i%2 == 1 {

			fmt.Println("buzz - ", i)
		}
		if i%2 == 0 {
			//this code doesn't work writes buzz and buzz and fizzbuzz on multiples of 5
			fmt.Println("fizz - ", i)
		}
		if i%5 == 0 {

			fmt.Println("fizzbuzz - ", i)
		}
		i++
	}

}
