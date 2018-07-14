package main   //works

import "fmt"

func main() {
	i := 0
	for i <= 10 {
		if i%2 == 0 {
			fmt.Println("even -", i)
		}else if i%2 == 1 {
			fmt.Println("odd -", i)
		}else {
			fmt.Println(i)
		}
		i++
	}
}
