package main

import "fmt"

const metertoyards float64 = 1.09361

func main() {
	var meters float64
	fmt.Print("Enter Meters swam:")
	fmt.Scan(&meters)
	yards := meters * metertoyards
	fmt.Println(meters, " meters is ", yards, " yards.")
}
