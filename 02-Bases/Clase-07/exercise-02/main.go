package main

import (
	"errors"
	"fmt"
)

var SalaryError = &InsufficientSalary{msg: "Error: salary is less than 10000"}

type InsufficientSalary struct {
	msg string
}

func (s *InsufficientSalary) Error() string {
	return s.msg
}

func checkSalary(salary int) error {
	if salary <= 10000 {
		return SalaryError
	}
	return nil
}

func main() {
	var salary int
	salary = 10001
	err := checkSalary(salary)
	if err != nil {
		if errors.Is(err, SalaryError) {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("Salary is valid")

}
