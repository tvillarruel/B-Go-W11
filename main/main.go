package main

import "fmt"

func main() {

	//-------------------------

	/* 

	// VAR, CONST AND FMT TESTS

	var name string
	var age int = 22
	const country string = "Argentina"

	fmt.Print("Hello, tell me your name please: ")
	fmt.Scan(&name)
	fmt.Println(name, " I know your age is: ", age, " and you live in: ", country) 
	*/
	 
	//-------------------------

	/*

	// MULTIPLE VAR AND CONSTS

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
	*/

	//-------------------------

	/* 

	// IF .. ELSE

	var (
		a = -1
		b = -1
		c = -1
		d = 1
	)

	if a < b && c < d {
		fmt.Println("Primer output")
	} else if a > b || c > d {
		fmt.Println("Segundo output")
	} else if d == 0 {
		fmt.Println("Tercer output")
	} else {
		fmt.Println("Cuarto output")
	}
	*/
	
	/*
	
	// SWITCH

	var op int

	fmt.Println("Ingresa un numero del 1 al 4")
	fmt.Scan(&op)
	
	switch op {
		case 1:
			fmt.Println("Uno")
		case 2:
			fmt.Println("Dos")
		case 3:
			fmt.Println("Tres")
		case 4:
			fmt.Println("Cuatro")
		default:
			fmt.Println("Cualquier otro numero")
	}
	*/

	fmt.Print("Hello World")

	
	 


}