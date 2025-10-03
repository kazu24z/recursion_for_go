/**
String invoiceNumber: 請求書番号
InvoiceItem invoiceItemHead: 購入したアイテムのリストの開始（連結リストの先頭）を表す InvoiceItem
double amountDue(bool taxes): 請求書の支払総額を計算します。InvoiceItemHead から始まるすべてのリスト項目を反復処理し、数量も考慮して計算する必要があります。入力が true に設定されている場合は、合計金額に、消費税分の 10% を加算してください。
*/

package main

type Invoice struct {
	InvoiceNumber   string
	InvoiceItemHead *InvoiceItem
}

func NewInvoice(number string, head *InvoiceItem) *Invoice {
	return &Invoice{
		InvoiceNumber:   number,
		InvoiceItemHead: head,
	}
}

func (i *Invoice) AmountDue(taxes bool) float64 {
	// iのInvoiceItemHeadからnext, nextって感じで辿っていく

	total := calcHelper(i.InvoiceItemHead)

	if taxes {
		total += total * 0.1
	}

	return total
}

// HEADを再帰的に処理するやつ
func calcHelper(node *InvoiceItem) float64 {
	// ベースケース
	if node == nil {
		return 0.0
	}

	// その回のnode分の料金を計算
	currentSum := node.Product.Price * float64(node.Quantity)

	// 次の回を呼び出す
	return currentSum + calcHelper(node.Next)
}
