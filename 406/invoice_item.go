package main

type InvoiceItem struct {
	Product  Product
	Quantity int
	Next     *InvoiceItem
}

func NewInvoiceItem(product Product, quantity int) *InvoiceItem {
	return &InvoiceItem{Product: product, Quantity: quantity}
}

func (i *InvoiceItem) GetTotalPrice() float64 {
	return i.Product.Price * float64(i.Quantity)
}
