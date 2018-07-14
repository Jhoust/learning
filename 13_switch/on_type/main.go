package main

import "fmt"

type contact struct {
	greeting string
	name     string
}

func SwitchOnType(x interface{}) {
	switch x.(type) { //x is flowing though the group switching on the new types it encounters the
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case contact:
		fmt.Println("contact")
	default:
		fmt.Println("unknown")
	}
}
func main() { //I guess the interfacing has to deal with this main func
	//As x moves through the cases and changes types  the interface is executing down here
	SwitchOnType(7) // or as the computer reads the main func it is moving the x through the type switch through the interface
	SwitchOnType("McLeod")
	var t = contact{"Good to see you,", "Tim"}
	SwitchOnType(t)
	SwitchOnType(t.greeting)
	SwitchOnType(t.name)
}
