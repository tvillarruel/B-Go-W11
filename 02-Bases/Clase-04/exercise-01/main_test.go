package main_test

import (
	"github.com/tvillarruel/GoTest/02-Bases/Clase-04/exercise-01/calculator"
	"testing"
	//"github.com/stretchr/testify/require"
)

func TestSalaryTaxCalculator_LessThan50k(t *testing.T) {
	//arrange

	//act
	salary := 49999.0
	expected := 0.0
	result := calculator.SalaryTaxCalculator(salary)

	//assert
	if result != expected {
		t.Errorf("Expected %f, got %f ", expected, result)
	}

}

func TestSalaryTaxCalculator_GreaterThan50kLessThan150k(t *testing.T) {
	//arrange

	//act
	salary := 149999.0
	expected := 25499.83
	result := calculator.SalaryTaxCalculator(salary)

	//assert
	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}

func TestSalaryTaxCalculator_GreaterThan150k(t *testing.T) {
	//arrange

	//act
	salary := 150001.0
	expected := 40500.27
	result := calculator.SalaryTaxCalculator(salary)

	//assert
	if result != expected {
		t.Errorf("Expected %.2f, got %.2f", expected, result)
	}
}
