package main

import "fmt"

func averagesCalculator(grades ...float64) float64 {

	var (
		avg   float64
		total float64
	)

	for i := 0; i < len(grades); i++ {
		total += grades[i]
	}
	avg = total / float64(len(grades))
	return avg
}

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
	average := averagesCalculator(notes...)
	fmt.Println("Notes: ", notes)
	fmt.Printf("Average: %.2f\n", average)

}
