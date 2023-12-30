package main

import "fmt"

func main(){
	/*
	Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. Según el siguiente mapa, ayudá a imprimir la edad de Benjamin.
	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	
	Por otro lado también es necesario:
	
	Saber cuántos de sus empleados son mayores de 21 años.
	Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
	Eliminar a Pedro del mapa.
	*/

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	fmt.Println("Benjamin's age: ", employees["Benjamin"])
	fmt.Println("Adding Federico...")
	employees["Federico"] = 25

	i:=0
	for _, age := range employees{
		if age > 21 {
			i++
		}
	}

	fmt.Println("Those over 21 years of age are: ", i)
	fmt.Println("Deleting Pedro...")
	delete(employees, "Pedro")
}