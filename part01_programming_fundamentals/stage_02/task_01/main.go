package main

import "fmt"

func main() {
	// const discountProduct = 10

	// titleProduct := "Клавиатура"
	// priceProduct := 2500.50
	// quantityProduct := 3
	// paidProduct := true

	// priceWithoutDiscount := priceProduct * float64(quantityProduct)
	// sizeDiscount := ((priceProduct * discountProduct) / 100) * float64(quantityProduct)
	// totalCost := priceWithoutDiscount - sizeDiscount
	// isSent := paidProduct && totalCost > 0

	// fmt.Println("Товар:", titleProduct)
	// fmt.Println("Количество:", quantityProduct)
	// fmt.Println("Стоимость без скидки:", priceWithoutDiscount)
	// fmt.Printf("Скидка: %.02f\n", sizeDiscount)
	// fmt.Println("К оплате:", totalCost)
	// fmt.Println("Можно отправить:", isSent)

	const discountPercent = 10
	productName := "Клавиатура"
	price := 2500.50
	quantity := 5
	isPaid := true

	subtotal := price * float64(quantity)
	discountAmount := subtotal * discountPercent / 100
	total := subtotal - discountAmount
	canBeSent := isPaid && total > 0

	fmt.Println("Товар:", productName)
	fmt.Println("Количество:", quantity)
	fmt.Printf("Стоимость без скидки: %.2f\n", subtotal)
	fmt.Printf("Скидка: %.2f\n", discountAmount)
	fmt.Printf("К оплате: %.2f\n", total)
	fmt.Println("Можно отправить:", canBeSent)

}
