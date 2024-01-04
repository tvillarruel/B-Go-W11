package calculator

func AveragesCalculator(grades ...float64) float64 {

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
