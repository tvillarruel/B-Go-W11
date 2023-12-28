package main

import "fmt"

func main() {

	// fmt.Print("Hello World")

	//-------------------------

	/* 
	VAR, CONST AND FMT TESTS

	var name string
	var age int = 22
	const country string = "Argentina"

	fmt.Print("Hello, tell me your name please: ")
	fmt.Scan(&name)
	fmt.Println(name, " I know your age is: ", age, " and you live in: ", country) 
	*/
	 
	//-------------------------

	var (
		firstName = "Thiago"
		lastName = "Villarruel"
		age = 22
		height = 1.58
		active = true
	)

	const (
		country = "Argentina"
		currentYear = 2023
	)

	fmt.Print("Hello, my full name is ", firstName, lastName," and I am ", age," years old. \n I measure ", height, ". \n Am I active?: ", active, "\n I live in ", country, " and today we are in ", currentYear)





	
	 


}