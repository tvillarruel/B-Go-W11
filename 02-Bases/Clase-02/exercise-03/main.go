package main

import "fmt"

func main(){

	//Elegí switch porque me resultó mas sencillo, podría haber usado un map y que el clave sea el numero y el valor el nombre del mes.
	//Tambien podria usar un while para no tener que ejecutar el programa cada vez que quiera probar una opcion nuevo
	
	var monthNumber int

	fmt.Print("Enter the number that corresponds to a month of the year: ")
	fmt.Scan(&monthNumber)

	switch monthNumber {
	case 1: fmt.Print("January")
	case 2: fmt.Print("February")
	case 3: fmt.Print("March")
	case 4: fmt.Print("April")
	case 5: fmt.Print("May")
	case 6: fmt.Print("June")
	case 7: fmt.Print("July")
	case 8: fmt.Print("August")
	case 9: fmt.Print("September")
	case 10: fmt.Print("October")
	case 11: fmt.Print("November")
	case 12: fmt.Print("December")
	default: fmt.Print("Wrong number")
	}
}