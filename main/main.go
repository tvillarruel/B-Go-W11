package main

import "fmt"

/*
const (
	Addition = "+"
	Substraction = "-"
	Multiplication = "*"
	Division = "/"
)

func calculate(a,b float64, operation string) float64 {
	switch operation {
	case Addition:
		 return a + b
	case Substraction:
		return a - b
	case Multiplication:
		return a * b
	case Division:
		if b != 0 {
			return a / b
		}
	}
	return 0
}
*/

func main() {

	//-------------------------

	/* 

	// VAR, CONST AND FMT TESTS

	var name string
	var age int = 22
	const country string = "Argentina"

	fmt.Print("Hello, tell me your name please: ")
	fmt.Scan(&name)
	fmt.Println(name, " I know your age is: ", age, " and you live in: ", country) 
	*/
	 
	//-------------------------

	/*

	// MULTIPLE VAR AND CONSTS

	var (
		firstName = "Thiago"
		lastName = "Villarruel"
		age = 22
		height = 1.58
		active = true
	)

	const (
		country = "Argentina"
		currentYear = 2023
	)

	fmt.Print("Hello, my full name is ", firstName, lastName," and I am ", age," years old. \n I measure ", height, ". \n Am I active?: ", active, "\n I live in ", country, " and today we are in ", currentYear)
	*/

	//-------------------------

	/* 

	// IF .. ELSE

	var (
		a = -1
		b = -1
		c = -1
		d = 1
	)

	if a < b && c < d {
		fmt.Println("Primer output")
	} else if a > b || c > d {
		fmt.Println("Segundo output")
	} else if d == 0 {
		fmt.Println("Tercer output")
	} else {
		fmt.Println("Cuarto output")
	}
	*/
	
	/*
	
	// SWITCH

	var op int

	fmt.Println("Ingresa un numero del 1 al 4")
	fmt.Scan(&op)
	
	switch op {
		case 1:
			fmt.Println("Uno")
		case 2:
			fmt.Println("Dos")
		case 3:
			fmt.Println("Tres")
		case 4:
			fmt.Println("Cuatro")
		default:
			fmt.Println("Cualquier otro numero")
	}
	*/

	/*

	// FOR

	for i := 0; i < 10; i++ {
		i++
		fmt.Println(i)
	}
 
	// WHILE

	var a int = 0
	for a < 10 {
		a++
		fmt.Print("Hola ")
	}
	*/
	
	/*

	// ARRAYS
	 var numbers [3]int
	 fruits := [3]string{"manzana","pera","uva"}
	 fmt.Println("Numeros: ", numbers)
	 fmt.Println("Frutas: ", fruits)

	 // SLICES VALOR POR DEFAULT ES UN PUNTERO HACIA UN ARRAY, (nil) 
	 // TRABAJA CON MEMORIA A MEDIDA QUE AVANZA EL PROGRAMA
	 names := make([]string,3)
	 names[0] = "Thiago"
	 names[1] = "Lucas"
	 names[2] = "Marcos"
	 fmt.Println("Names: ", names)

	 names = append(names, "Jose", "Jorge")
	 fmt.Println("Names: ", names)

	 // MAP
	 // MAKE() para array no sirve, el make permite en slices y map 
	 // inicializarlos y poder especificar longitud y capacidad
	 // (capacidad es opcional especificar, si no se especifica
	 // va a ser igual a la longitud)
	 ages := make(map[string]uint)
	 ages["Thiago"] = 22
	 age := ages["Thiago"]
	 fmt.Println(age)


	 // RANGE es como un foreach, va iterando ¿segun el tipo?

	 */
	
	/*
	fmt.Println("Suma: ", calculate(6,2,Addition))
	fmt.Println("Resta: ", calculate(6,2,Substraction))
	fmt.Println("Multiplicacion: ", calculate(6,2,Multiplication))
	fmt.Println("Division: ", calculate(6,2,Division))
	*/
	fmt.Print("")
}