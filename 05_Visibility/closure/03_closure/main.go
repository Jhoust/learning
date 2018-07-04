package main

import "fmt"

func wrapper() func() int {
	x := 0					//the break down is that x is declared in top of the wrapper function
								//func wrapper is called at the bottom for printing
								//the next anon func after the wrapper function is what sets things up so it's needed


	return func() int {
		//x := 3

			//if x is defined at this level it gets added once and prints twice, but it does not get printed as it's assigned value
			// so it shows its value +1 twice, which defeats the purpose of the function calling the return of x first and then adding another function to increase it by 1

										//it's going to return a value which happens to be another function adding up x++
			x++
										//which then finally returns x
				return x
			}

	}
}

func main() {

	increment := wrapper ()     //here func wrapper gets a new name that accurately describes it's inner workings our old friend (increment)

	fmt.Println(increment())    //(increment is then printed out. Don't forget to add the () it's a function and just just a word!
	fmt.Println(increment())
}