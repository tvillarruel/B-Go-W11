package main

import "fmt"

func main(){

	var (
		word string
		letters int
	)

	fmt.Print("Write a word: ")
	fmt.Scan(&word)

	letters = len(word)

	fmt.Println("The word is", word, ", has", letters, "letters");

	for i:= 0; i < len(word); i++ {
		fmt.Print(string(word[i]), " ")
		// Funcion string se utiliza para convertir un valor tipo rune (caracter unicode) a su representacion en cadena de texto.
	}


}