package main

import "fmt"

type Product interface {
	Price() float64
}

type Small struct {
	PriceValue float64
}

func (sp *Small) Price() float64 {
	return sp.PriceValue
}

type Medium struct {
	PriceValue float64
}

func (mp *Medium) Price() float64 {
	return mp.PriceValue *1.03
}

type Large struct {
	PriceValue float64
}

func (lp *Large) Price() float64 {
	return (lp.PriceValue * 1.06) + 2500
}

func factory(productType string, price float64) Product {
	switch productType {
	case "Small":
		return &Small{price}
	case "Medium":
		return &Medium{price}
	case "Large":
		return &Large{price}
	default:
		return nil
	}
}

func main() {
	
	product1 := factory("Small", 999.99)
	product2 := factory("Medium", 999.99)
	product3 := factory("Large", 999.99)

	fmt.Println("Small:", product1.Price())
	fmt.Println("Medium:", product2.Price())
	fmt.Println("Large:", product3.Price())
}