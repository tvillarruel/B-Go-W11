package main

import (
	"fmt"
	"github.com/tvillarruel/GoTest/02-Bases/Clase-04/exercise-02/calculator"
)

func main() {

	notes := make([]float64, 0)

	fmt.Print("Enter notes separated by spaces (enter a negative note to end):")
	for {
		var note float64
		fmt.Scan(&note)
		if note < 0 {
			break
		}
		notes = append(notes, note)
	}
	average := calculator.AveragesCalculator(notes...)
	fmt.Println("Notes: ", notes)
	fmt.Printf("Average: %.2f\n", average)

}
