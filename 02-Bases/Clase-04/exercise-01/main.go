package main

import (
	"fmt"
	"github.com/tvillarruel/GoTest/02-Bases/Clase-04/exercise-01/calculator"
)

func main(){

	var salary float64

	fmt.Print("Tell me your salary: ")
	fmt.Scan(&salary)
	
	result := fmt.Sprintf("%.2f", calculator.SalaryTaxCalculator(salary))
	fmt.Println("Tax applied:", result)
}