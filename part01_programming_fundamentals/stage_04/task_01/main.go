package main

import "fmt"

func main() {
	discountPercent := 10
	// discountPercent := -5
	// discountPercent := 120
	if !isValidDiscount(discountPercent) {
		fmt.Println("Некорректная скидка")
		return
	}
}

func isValidDiscount(discount int) bool {
	return discount >= 0 && discount <= 100
}
