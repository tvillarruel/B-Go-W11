package main

import (
	"errors"
	"fmt"
)

var SalaryError = errors.New("Error: salary is less than 10000")

type InsufficientSalaryError struct {
	msg error
}

func (s *InsufficientSalaryError) method() error {
	s.msg = errors.New("Error: Error: salary is less than 10000")
	return s.msg
}

func checkSalary(salary int) error {
	if salary <= 10000 {
		return SalaryError
	}
	return nil
}

func main() {
	var salary = 10
	err := checkSalary(salary)
	var msg = InsufficientSalaryError{
		msg: errors.New(""),
	}

	if err != nil {

		if errors.Is(err, SalaryError) {
			fmt.Println(msg.method())
			return
		}
	}
	fmt.Println("Salary is valid")
}
