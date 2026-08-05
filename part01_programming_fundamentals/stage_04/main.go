package main

import "fmt"

func main() {
	price := 5000
	discountPercent := 10

	discountAmount := calculateDiscount(price, discountPercent)
	finalPrice := calculateFinalPrice(price, discountAmount)

	printOrderSummary(price, discountAmount, finalPrice)

}

func calculateFinalPrice(price, discountAmount int) int {
	return price - discountAmount
}

func printOrderSummary(price, discountAmount, finalPrice int) {
	fmt.Println("Цена:", price)
	fmt.Println("Скидка:", discountAmount)
	fmt.Println("Итого:", finalPrice)
}

func calculateDiscount(price, discount int) int {
	return price * discount / 100
}

func getAccessMessage(age int) string {

	switch {
	case age < 0:
		return "Некорректный возраст"
	case age < 18:
		return "Доступ запрещён"
	default:
		return "Доступ разрешён"
	}
}

func multiply(a, b int) int {
	product := a * b
	return product
}

func divideNumbers(a, b int) (int, int, bool) {
	if b == 0 {
		return 0, 0, false
	}
	return a / b, a % b, true
}

func calculatePrice(price, discout int) (int, int) {
	discountAmount := price * discout / 100
	finalPrice := price - discountAmount

	return discountAmount, finalPrice
}

func calculateTotal(price int, quantity int) int {
	return price * quantity
}

func printProduct(name string, price int) {
	fmt.Println("Товар:", name)
	fmt.Println("Цена:", price)
}

func printUser(name string, age int) {
	fmt.Println("Пользователь:", name)
	fmt.Println("Возраст:", age)
}

func printMessage(message string) {
	fmt.Println(message)
}

func printStartMessage() {
	fmt.Println("Программа запущена")
}

func printSeparator() {
	fmt.Println("----------")
}
