package main_test

import (
	"github.com/tvillarruel/GoTest/02-Bases/Clase-04/exercise-02/calculator"
	"testing"
)

func TestAveragesCalculator(t *testing.T) {
	//arrange

	//act
	grades := []float64{10.0, 10.0, 1.0}
	expected := 7.0
	result := calculator.AveragesCalculator(grades...)

	//assert
	if result != expected {
		t.Errorf("Expected %f, got %f", expected, result)
	}
}
