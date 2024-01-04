package main

import "fmt"

func main() {

	var (
		lastName string = "Smith"
		//var age int = "35"
		age int = 35
		//var salary string = 45857.90
		salary    float64 = 45857.90
		firstName string  = "Mary"
	)

	// boolean := "false"
	boolean := false

	fmt.Println(lastName, "\n", age, "\n", salary, "\n", firstName, "\n", boolean)
}
