package main

import (
	"fmt"
)

type Product struct {
	title string
	price float64
}

func NewProduct(title string, price float64) Product {
	return Product{
		title: title,
		price: price,
	}
}
func main() {
	product1 := NewProduct("shampoo", 10)
	product2 := NewProduct("conditioner", 5)

	// テストケースの出力を確認
	fmt.Println(product1.title)
	fmt.Println(product1.price)
	fmt.Println(product2.title)
	fmt.Println(product2.price)

}
