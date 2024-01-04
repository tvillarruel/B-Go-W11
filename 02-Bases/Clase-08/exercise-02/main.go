package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// fmt.Println(os.Getwd()) muestra la ruta en la que esta trabajando
	defer func() {
		fmt.Println("ejecucion finalizada")
	}()

	file, err := os.Open("/Users/tvillarruel/Desktop/GoTest/02-Bases/Clase-08/exercise-02/customer.txt")
	if err != nil {
		panic("The indicated file was not found or is damaged")
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	data := string(bytes)
	fmt.Println(data)
}
