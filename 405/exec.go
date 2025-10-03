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

	// テスト
	fmt.Println(firstItem.GetTotalPrice())
	fmt.Println(secondItem.GetTotalPrice())
	fmt.Println(firstItem.Next.GetTotalPrice())
	fmt.Println(firstItem.Next.Next.GetTotalPrice())
	fmt.Println(firstItem.Next.Product.Price)
	fmt.Println(firstItem.Next.Next.Product.Title)
}

/**
Product をまとめて購入し、請求書に入れるため、請求書 1 つ 1 つの項目が必要になります。以下に従って、InvoiceItem クラスを作成し、テストケースを出力してください。

Product product: 製品オブジェクト
int quantity: 購入する製品の数
InvoiceItem next: 請求書の次の項目
double getTotalPrice(): 購入する数量に基づいて、製品の合計価格を計算します。

*/
