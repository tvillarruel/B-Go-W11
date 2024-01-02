package main

import "fmt"

func salaryCalculator(minPerMonth float64, category string) (salary float64) {
	
	var totalHours float64 = minPerMonth/60

	switch category { 
	case "C":
		salary = 1000 * totalHours
		return
	case "B":
		salary = 1500 * totalHours
		salary += (0.20*salary)
		return
	case "A":
		salary = 3000 * totalHours
		salary += (0.50*salary)
		return
	default: fmt.Println("Error: wrong category")
	}
	return 0
}

func main(){

	var (
		minPerMonth float64
		category string
	)

	fmt.Print("Enter your category (A, B or C): ")
	fmt.Scanln(&category)
	fmt.Print("Enter the number of minutes worked: ")
	fmt.Scanln(&minPerMonth)

	result := salaryCalculator(minPerMonth, category)
	fmt.Println("Your salary is: ", result)
}