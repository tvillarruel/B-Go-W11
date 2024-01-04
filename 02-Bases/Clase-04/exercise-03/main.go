package main

import (
	"fmt"
	"github.com/tvillarruel/GoTest/02-Bases/Clase-04/exercise-03/calculator"
)

func main() {

	var (
		minPerMonth float64
		category    string
	)

	fmt.Print("Enter your category (A, B or C): ")
	fmt.Scanln(&category)
	fmt.Print("Enter the number of minutes worked: ")
	fmt.Scanln(&minPerMonth)

	result, err := calculator.SalaryCalculator(minPerMonth, category)
	if err != nil {
		fmt.Print("Unexpected error: ", err)
	} else {
		fmt.Println("Your salary is: ", result)
	}
}
