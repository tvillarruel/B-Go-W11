package main

import (
	"errors"
	"fmt"
)

func salaryByHours(numHours float64, valuePerHour float64) (float64, error) {

	var salary float64 = numHours * valuePerHour

	if numHours < 80 || numHours < 0 {
		return 0, errors.New("Error: the worker cannot have worked less than 80 hours per month\n")
	} else if salary >= 150000 {
		discount := salary * 0.10
		salary -= discount
		return salary, nil
	}
	return salary, nil
}

func main() {

	result, err := salaryByHours(80, 2000)
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Printf("Salary: %.2f\n", result)
}
