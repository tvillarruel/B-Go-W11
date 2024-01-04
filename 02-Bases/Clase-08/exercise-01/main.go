package main

import (
	"fmt"
	"os"
)

func main() {

	defer func() {
		fmt.Println("ejecucion finalizada")
	}()
	_, err := os.Open("customer.txt")
	if err != nil {
		panic("The indicated file was not found or is damaged")

	}

}
