package main

import "fmt"

func main(){

	var (
		age int
		active bool
		seniority int
		salary float64

	)

	fmt.Println("Complete the data in order with a space beetwen answers: \n 1) Age: \n 2) Are you active? (true or false): \n 3) Seniority (in years): \n 4) Salary: ")
	fmt.Scan(&age, &active, &seniority, &salary)

	if age < 22 || active == false || seniority <= 1 {
		fmt.Println("You do not meet some of the requeriments to receive a loan: \n- Be over 22 years old\n- Be employed\n- Have been in the company for one year")
	} else if salary <= 100000 {
		fmt.Println("You will be granted a loan but you will be charged interest because your salary is less than 100,000 ARS")
	} else {
		fmt.Println("You will be granted a loan and no interest will be charged")
	}
}