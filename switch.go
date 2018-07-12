package main

import "fmt"

var a1 string

func main() {
	var i = 0
	for
	{ i++
		if i%2 == 0 {
			a1 = "a"
		}
			if i%2 == 1 {
				a1 = "b"
			}
				if i%5 == 0 {
					a1 = "c"
				}
					if i >= 11 {
						break
					}
					switch a1 {
					case "a":
						fmt.Println("buzz")
					case "b":
						fmt.Println("fizz")
					case "c":
						fmt.Println("fizzbuzz")
					default:
						fmt.Println("you broke it")
					}
				}
			}
