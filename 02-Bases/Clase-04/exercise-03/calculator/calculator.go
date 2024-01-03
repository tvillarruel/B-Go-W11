package calculator

import "fmt"

func SalaryCalculator(minPerMonth float64, category string) (salary float64, err error) {
	
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
	default: 
		fmt.Println("Error: wrong category")
	}
	return 0, err
}