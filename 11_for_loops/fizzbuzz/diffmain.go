package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 1{
			fmt.Println(i)
			fmt.Println("buzz")
		}
			if i%2 == 0 {
				fmt.Println(i)								//this code doesn't work writes buzz and buzz and fizzbuzz on multiples of 5
				fmt.Println("fizz")
			}
				if i%5 == 0 {
					fmt.Println(i)
					fmt.Println("fizzbuzz")
				}

			}

		}


