package main

import "fmt"

func salaryTaxCalculator(salary float64) float64 {
	switch {
	case salary > 50000 && salary < 150000:
		salary *= 0.17
		return salary

	case salary > 150000:
		salary *= 0.27
		return salary
	}
	return 0

}

func main() {

	var salary float64

	fmt.Print("Tell me your salary: ")
	fmt.Scan(&salary)

	result := fmt.Sprintf("%.2f", salaryTaxCalculator(salary))
	fmt.Println("Tax applied:", result)
}
