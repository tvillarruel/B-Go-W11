package main

import "fmt"

type Product struct {
	ID int
	Name string
	Price float64
	Description string
	Category string
}

var Products = []Product {
	{
		ID: 1, 
		Name: "Laptop", 
		Price: 1499.99,
		Description: "A very good laptop",
		Category: "Electronics",
	},
	{
		ID: 2,
		Name: "Smartphone", 
		Price: 999.99,
		Description: "A very good smartphone",
		Category: "Electronics",
		
	},
	{
		ID: 3, 
		Name: "Book", 
		Price: 1499.99,
		Description: "A very good books",
		Category: "Books",
	},
}

func (p Product) Save(){
	Products = append(Products, p)
	fmt.Println("Product saved")

}

func (p Product) GetAll(){
	fmt.Println("All products:")
	for _, product := range Products {
		fmt.Printf("ID: %d\nName: %s\nPrice: %.2f\nDescription: %s\nCategory: %s\n", product.ID, product.Name, product.Price, product.Description, product.Category)
	}
}

func getById(id int) (product Product) {
	for _, product := range Products{
		if product.ID == id {
			return product
		}
	}
	return product
}

func main(){

	var newProduct = Product {
		ID: 4,
		Name: "Car",
		Price: 9999.99,
		Description: "A very good car",
		Category: "Vehicles",
	}

	newProduct.Save()

	newProduct.GetAll()

	id := 4
	productByID := getById(id)
	
	fmt.Printf("Product with id %d: %s", id, productByID.Name)
}

