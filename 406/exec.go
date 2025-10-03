package main

import "fmt"

func main() {
	// 購入するProduct インスタンス
	product1 := NewProduct("shampoo", 10)
	product2 := NewProduct("conditioner", 5)
	product3 := NewProduct("tooth brush", 3)

	// 請求書インスタンス
	firstItem := NewInvoiceItem(product1, 7)
	secondItem := NewInvoiceItem(product2, 9)
	thirdItem := NewInvoiceItem(product3, 10)

	// 1番からnextに格納
	firstItem.Next = secondItem
	secondItem.Next = thirdItem

	invoice := NewInvoice("UC1234567890", firstItem)

	// テスト
	fmt.Println(invoice.AmountDue(false))
	fmt.Println(invoice.AmountDue(true))
}
