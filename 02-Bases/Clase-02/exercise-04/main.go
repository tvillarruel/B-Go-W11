package main

import "fmt"

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println("Benjamin's age: ", employees["Benjamin"])
	fmt.Println("Adding Federico...")
	employees["Federico"] = 25

	i := 0
	for _, age := range employees {
		if age > 21 {
			i++
		}
	}

	fmt.Println("Those over 21 years of age are: ", i)
	fmt.Println("Deleting Pedro...")
	delete(employees, "Pedro")
}
