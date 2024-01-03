package main

import "fmt"

type Person struct {
	ID int
	Name string
	DateOfBirth string
}

type Employee struct {
	ID int
	Position string
	Person Person
}

func (emp Employee) PrintEmployee() {
	fmt.Printf("Employee ID: %d\nEmployee position: %s\nPerson ID: %d\nPerson name: %s\nPerson date of birth: %s\n", emp.ID, emp.Position, emp.Person.ID, emp.Person.Name, emp.Person.DateOfBirth)
}

func main(){

	var p1 = Person {
		ID: 1,
		Name: "Thiago",
		DateOfBirth: "26/10/2001",
	}

	var emp1 = Employee {
		ID: 1,
		Position: "Software Developer",
		Person: p1,
	}
	

	emp1.PrintEmployee()
}