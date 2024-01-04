package main

import "fmt"

var SalaryError = &InsufficcientSalaryError{"Error: the salary entered dos not reach taxable minimum"}

type InsufficcientSalaryError struct {
	msg string
}

func (s *InsufficcientSalaryError) Error() string {
	return s.msg
} 

func checkSalary(salary int) error {
	if salary < 150000 {
		return SalaryError
	}
	return nil
}

func main() {

	var salary int

	salary = 150001

	err := checkSalary(salary)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Must pay tax")
}