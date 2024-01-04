package main_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tvillarruel/GoTest/02-Bases/Clase-04/exercise-03/calculator"
	"testing"
)

func TestSalaryCalculator_CategoryC(t *testing.T) {
	//arrange

	//act
	minPerMonth := 60.0
	category := "C"
	expected := 1000.0
	result, err := calculator.SalaryCalculator(minPerMonth, category)

	//assert
	//assert
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}
	if expected != result {
		t.Fatalf("Expected %f, got %f", expected, result)
	}
}

func TestSalaryCalculator_CategoryB(t *testing.T) {
	//arrange

	//act
	minPerMonth := 60.0
	category := "B"
	expected := 1800.0
	result, err := calculator.SalaryCalculator(minPerMonth, category)

	//assert
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}
	if expected != result {
		t.Fatalf("Expected %f, got %f", expected, result)
	}
}

func TestSalaryCalculator_CategoryA(t *testing.T) {
	//arrange

	//act
	minPerMonth := 60.0
	category := "A"
	expected := 4500.0
	result, err := calculator.SalaryCalculator(minPerMonth, category)

	//assert
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}
	require.Equal(t, expected, result, "Verify the expected result is equal to the actual result")
}
