package main

import "fmt"

func multiply() func() float64 {
	 c := 1.0
	 f := 0.0
	return func() float64 {
		f = c*1.8 + 32

		return f
	}
}


func main(){
	temp := multiply()
	fmt.Println(temp())
}