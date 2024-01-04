package main

import "fmt"

func animal(animalType string) (func(int) float64, string) {
	switch animalType {
	case "dog":
		return animalDog, "Dog selected"
	case "cat":
		return animalCat, "Cat selected"
	case "hamster":
		return animalHamster, "Hamster selected"
	case "tarantula":
		return animalTarantula, "Tarantula selected"
	default:
		return nil, "Wrong animal selected"
	}

}

func animalDog(numDogs int) float64 {
	var food int = 10
	return float64(numDogs) * float64(food)
}

func animalCat(numCat int) float64 {
	var food int = 5
	return float64(numCat) * float64(food)
}

func animalHamster(numHamster int) float64 {
	var food float64 = 0.250
	return float64(numHamster) * float64(food)
}

func animalTarantula(numTarantula int) float64 {
	var food float64 = 0.150
	return float64(numTarantula) * float64(food)
}

func main() {

	const (
		dog       = "dog"
		cat       = "cat"
		hamster   = "hamster"
		tarantula = "tarantula"
	)

	typeAnimalFunc, msg := animal(cat)

	if typeAnimalFunc != nil {
		result := typeAnimalFunc(10)
		fmt.Printf("%s, %.2f Kg of food\n", msg, result)
	} else {
		fmt.Println(msg)
	}

}
