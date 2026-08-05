package main

import "fmt"

func main() {
	// var prices [3]int

	// prices[0] = 500
	// prices[1] = 1000
	// prices[2] = 300
	// fmt.Println(prices)
	// fmt.Println("Первый товар:", prices[0])
	// fmt.Println("Последний товар:", prices[2])

	// prices := [3]int{500, 1200, 300}
	// // prices := [...]int{500, 1200, 300}

	// fmt.Println("Количество товаров:", len(prices))

	// prices := [...]int{500, 1200, 300}

	// for i := 0; i < len(prices); i++ {
	// 	fmt.Println("Индекс:", i, "Цена:", prices[i])
	// }

	// grades := [...]int{5, 4, 3, 5, 4}
	// grades[2] = 5

	// fmt.Println(grades)

	// for i := 0; i < len(grades); i++ {
	// 	// fmt.Println("Оценка", i+1, ":", grades[i])
	// 	fmt.Printf("Оценка %d: %d\n", i+1, grades[i])
	// }

	// grades := [...]int{5, 4, 5, 5, 4}

	// // for index, grade := range grades {
	// // 	fmt.Printf("Оценка %d: %d\n", index+1, grade)
	// // }

	// for _, grade := range grades {
	// 	fmt.Println(grade)
	// }

	// prices := [...]int{500, 1200, 300, 800}

	// totalPrice := 0

	// for index, value := range prices {
	// 	totalPrice += value
	// 	fmt.Printf("Цена %d: %d\n", index+1, value)
	// }

	// fmt.Println("Общая сумма:", totalPrice)

	// tasks := []string{"Изучить массивы", "Изучить срезы"}

	// fmt.Println(tasks)
	// fmt.Println("Количество задач:", len(tasks))

	// tasks[1] = "Освоить срезы"
	// fmt.Println(tasks)

	// tasks = append(tasks, "Изучить maps")
	// tasks = append(tasks, "Написать практику", "Повторить материал")
	// fmt.Println(tasks)

	// products := []string{"Клавиатура", "Мышь"}
	// products[1] = "Беспроводная мышь"
	// products = append(products, "Монитор")
	// products = append(products, "Колонки", "Наушники")

	// for index, value := range products {
	// 	fmt.Printf("%d. %s\n", index+1, value)
	// }

	// fmt.Println("Всего товаров:", len(products))

	// products := []string{
	// 	"Клавиатура",
	// 	"Беспроводная мышь",
	// 	"Монитор",
	// 	"Колонки",
	// 	"Наушники",
	// }

	// printSearchResult(products, "Монитор")
	// printSearchResult(products, "Ноутбук")

	// products = append(products[:2], products[3:]...)

	// fmt.Println(products)
	// products = removeProduct(products, "Монитор")
	// fmt.Println(products)

	prices := map[string]int{
		"Клавиатура": 5000,
		"Мышь":       2500,
		"Монитор":    18000,
	}

	// fmt.Println(prices)
	// fmt.Println("Цена клавиатуры:", prices["Клавиатура"])

	prices["Наушники"] = 4000
	prices["Мышь"] = 3000
	prices["Колонки"] = 6000
	prices["Монитор"] = 17500

	// fmt.Printf("Цена мышки: %d\n", prices["Мышь"])

	// fmt.Printf("Количество товаров: %d\n", len(prices))

	// printProductPrice(prices, "Монитор")
	// printProductPrice(prices, "Ноутбук")

	removeProductPrice(prices, "Мышь")
	removeProductPrice(prices, "Ноутбук")

	fmt.Println("Количество товаров:", len(prices))

	printCatalog(prices)

	updatePriductPrice(prices, "Монитор", 20000)
	updatePriductPrice(prices, "Ноутбук", 50000)

	printProductPrice(prices, "Монитор")

	addProduct(prices, "Веб-камера", 4500)
	addProduct(prices, "Монитор", 25000)
}

func addProduct(prices map[string]int, product string, price int) {
	if _, exists := prices[product]; exists {
		fmt.Printf("%s уже существует\n", product)
		return
	}
	prices[product] = price
	fmt.Printf("Товар %s добавлен.\n", product)
}

func printCatalog(prices map[string]int) {
	totalPrice := 0
	for key, value := range prices {
		totalPrice += value
		fmt.Printf("%s: %d\n", key, value)
	}
	fmt.Printf("Общая стоимость: %d\n", totalPrice)
}

func updatePriductPrice(prices map[string]int, product string, newPrice int) {
	if _, exists := prices[product]; !exists {
		fmt.Printf("%s не найден\n", product)
		return
	}

	prices[product] = newPrice
	fmt.Printf("Цена товара %s изменена на %d\n", product, newPrice)
}

func removeProductPrice(prices map[string]int, product string) {
	_, exists := prices[product]
	if !exists {
		fmt.Printf("%s не найден\n", product)
		return
	}
	delete(prices, product)
	fmt.Printf("Товар %s удалён\n", product)
}

func printProductPrice(prices map[string]int, product string) {
	price, exists := prices[product]
	if !exists {
		fmt.Printf("%s не найден\n", product)
		return
	}
	fmt.Printf("%s: %d\n", product, price)
}

func removeProduct(products []string, target string) []string {
	index := findProduct(products, target)
	if index == -1 {
		return products
	}
	return append(products[:index], products[index+1:]...)
}

func printSearchResult(products []string, target string) {
	foundIndex := findProduct(products, target)

	if foundIndex == -1 {
		fmt.Printf("%s не найден\n", target)
		return
	}
	fmt.Printf("%s найден под номером %d\n", target, foundIndex+1)
}

func findProduct(products []string, target string) int {
	for index, product := range products {
		if product == target {
			return index
		}
	}
	return -1
}
