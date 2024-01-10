package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

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

func main() {

	//* Open file
	products, err := openFile("/Users/tvillarruel/Desktop/GoTest/03-Web/Clase-02/exercise-02/products.json")
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
	})
	//* Open server
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Println("Error:", err)
		return
	}
}
