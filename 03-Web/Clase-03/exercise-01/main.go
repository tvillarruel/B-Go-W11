package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
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

func openFile(fileName string) ([]byte, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	return file, nil

}

func fillSlice(data []byte) ([]Product, error) {
	err := json.Unmarshal(data, &Products)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	return Products, nil
}

func getPing() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	}
}

func getAll(products []Product) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	}

}

func getById(products []Product) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		for _, product := range products {
			if strconv.Itoa(product.Id) == id {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(product)
			}
		}

	}
}

func getByPrice(products []Product) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		priceGt := r.URL.Query().Get("priceGt")
		if priceGt == "" {
			http.Error(w, "Missing priceGt parameter", http.StatusBadRequest)
			return
		}

		priceThreshold, err := strconv.ParseFloat(priceGt, 64)
		if err != nil {
			http.Error(w, "Invalid priceGt value", http.StatusBadGateway)
			return
		}

		var filteredProducts []Product
		for _, products := range products {
			if products.Price > priceThreshold {
				filteredProducts = append(filteredProducts, products)
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(filteredProducts)
	}
}

// ! DECODER DE JSON A GO
// ! ENCODER DE GO A JSON
func create(products []Product) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var newProduct Product
		if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte("Invalid body"))
			return
		}

		err := validateProduct(newProduct)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		newProduct.Id = inferID(products)

		Products = append(Products, newProduct)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newProduct)
	}
}

func validateProduct(product Product) error {
	if product.Name == "" || product.Quantity == 0 || product.CodeValue == "" || product.Expiration == "" || product.Price == 0.0 {
		return errors.New("there cannot be empty fields")
	}

	for _, p := range Products {
		if p.CodeValue == product.CodeValue {
			return errors.New("code value must be unique")
		}
	}

	_, err := time.Parse("02/01/2006", product.Expiration)
	if err != nil {
		return errors.New("invalid date format for Expiration")
	}
	return nil
}

func inferID(products []Product) int {
	highestID := 0
	for _, p := range products {
		if p.Id > highestID {
			highestID = p.Id
		}
	}
	return highestID + 1
}
func main() {

	//* Open file
	products, err := openFile("/Users/tvillarruel/Desktop/GoTest/03-Web/Clase-03/exercise-01/products.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//* Load data on slice
	productSlice, err := fillSlice(products)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//* New chi router
	r := chi.NewRouter()

	//*ENDPOINT PING
	r.Get("/ping", getPing())
	//* ENDPOINTS GROUP
	r.Route("/products", func(r chi.Router) {
		r.Get("/", getAll(productSlice))
		r.Get("/{id}", getById(productSlice))
		r.Get("/search", getByPrice(productSlice))
		r.Post("/", create(productSlice))
	})
	//* Open server
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("Error:", err)
		return
	}
}
