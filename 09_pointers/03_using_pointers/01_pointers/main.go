package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a) //43
	fmt.Println(&a)    //memory address

	var b *int = &a  //memory address given to b
	fmt.Println(b)   //b prints memory address
	fmt.Println(*b)   // b dereferences with ^ and prints 43 from top

	*b = 42           //^b changes address to the value of 42
	fmt.Println(a)      // a is still pointing to the same address which is has been given the new value of 42 prints 42
}
/*this is useful
we can pass a memory address instead of
a bunch of values (we cam pass a reference)
and then we can still change the value of whatever is
stored at that memory address
this makes our programs more performant
we don't have to pass around large amounts of data
we only have to pass around addresses
 */