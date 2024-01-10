package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

var Products []Product

func OpenFile(fileName string) ([]byte, error) {
	//* Open file
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	return file, nil

}
func fillSlice(file []byte) ([]Product, error) {
	err := json.Unmarshal(file, &Products)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	}
	return Products, nil
}

func main() {

	data, err := OpenFile("/Users/tvillarruel/Desktop/GoTest/03-Web/Clase-02/exercise-01/products.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	products, err := fillSlice(data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(products)
}
