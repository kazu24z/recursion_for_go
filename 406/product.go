package main

type Product struct {
	Title string
	Price float64
}

func NewProduct(title string, price float64) Product {
	return Product{
		Title: title,
		Price: price,
	}
}
