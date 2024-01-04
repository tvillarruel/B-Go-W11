package main

import "fmt"

type Alumnos struct {
	Name string
	Apellido string
	DNI int
	Fecha string
}

var AlumnoList = []Alumnos {}

func (a Alumnos) save() {
	AlumnoList = append(AlumnoList, a)
	fmt.Println("Alumno saved")
}

func (a *Alumnos) detalle(){
	for _, alumno := range AlumnoList {
		fmt.Printf("Name: %s\nApellido: %s\nDNI: %d\nFecha: %s\n", alumno.Name, alumno.Apellido, alumno.DNI, alumno.Fecha)
	}
}


func main() {

	var newAlumno = Alumnos {
		Name: "Thiago",
		Apellido: "Villarruel",
		DNI: 12345678,
		Fecha: "20/12/2023",



	}

	var newAlumno2 = Alumnos {
		Name: "Marcos",
		Apellido: "Gomez",
		DNI: 87654321,
		Fecha: "17/3/1995",


		
	}

	newAlumno.save()
	newAlumno2.save()
	newAlumno.detalle()
	
}