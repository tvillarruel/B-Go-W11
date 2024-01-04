package main

import "fmt"

func calcType(calcType string) (func(...int) float64, string) {
	switch calcType {
	case "minimum":
		return calcMin, "Selected minimum calculation"
	case "maximum":
		return calcMax, "Selected maximum calculation"
	case "average":
		return calcAvg, "Selected average calculation"
	default:
		return nil, "Calculation type not supported"
	}
}

func calcMin(grades ...int) float64 {
	var min int
	min = grades[0]
	for i := 0; i < len(grades); i++ {
		if min > grades[i] {
			min = grades[i]
		}
	}
	return float64(min)
}

func calcMax(grades ...int) float64 {
	var max int
	max = grades[0]
	for _, note := range grades {
		if max < note {
			max = note
		}
	}
	return float64(max)
}

func calcAvg(grades ...int) float64 {
	var sum int
	for _, note := range grades {
		sum += note
	}
	return float64(sum) / float64(len(grades))
}

func main() {

	grades := []int{90, 86, 44, 34, -23, 77, 9, 2, 1, -18, 0}

	const (
		minimum = "minimum"
		average = "average"
		maximum = "maximum"
	)

	calcFunc, msg := calcType(minimum)

	if calcFunc != nil {
		result := calcFunc(grades...)
		fmt.Printf("%s: %.2f\n", msg, result)
	} else {
		fmt.Println(msg)
	}

}
