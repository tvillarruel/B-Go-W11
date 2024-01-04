package calculator

func SalaryTaxCalculator(salary float64) float64 {
	switch {
	case salary > 50000 && salary < 150000:
		salary *= 0.17
		return salary

	case salary > 150000:
		salary *= 0.27
		return salary
	}
	return 0

}
