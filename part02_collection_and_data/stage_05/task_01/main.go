package main

import "fmt"

func main() {
	catalog := make(map[string]int)

	addProduct(catalog, "i9-10850", 9800)
	addProduct(catalog, "i9-10900", 10300)
	addProduct(catalog, "i3-10100", 2350)
	addProduct(catalog, "i3-10100", 2350)
	printCatalog(catalog)
	searchProduct(catalog, "i3-10100")
	searchProduct(catalog, "i3-10200")
	changePrice(catalog, "i9-10850", 9500)
	removeProduct(catalog, "i9-10900")
	printCatalog(catalog)
}

func removeProduct(catalog map[string]int, target string) {
	if _, exists := catalog[target]; !exists {
		fmt.Printf("Товар %s не найден.\n", target)
		return
	}

	delete(catalog, target)
	fmt.Printf("Продукт %s - удалён\n", target)
}

func changePrice(catalog map[string]int, target string, newPrice int) {
	if _, exists := catalog[target]; !exists {
		fmt.Printf("Товар %s не найден.\n", target)
		return
	}
	catalog[target] = newPrice
	fmt.Printf("Товар %s,цена изменена на %d.\n", target, newPrice)

}

func searchProduct(catalog map[string]int, target string) {
	if _, exists := catalog[target]; !exists {
		fmt.Printf("Товар %s не найден.\n", target)
		return
	}
	fmt.Printf("Товар %s найден.\n", target)
}

func printCatalog(catalog map[string]int) {
	count := 0
	for key, value := range catalog {
		count++
		fmt.Printf("%d. %s: %d\nВсего товаров: %d\n", count, key, value, len(catalog))
	}
}

func addProduct(catalog map[string]int, product string, price int) {
	if _, exists := catalog[product]; exists {
		fmt.Printf("Продукст %s уже существует!\n", product)
		return
	}

	catalog[product] = price
	fmt.Printf("Продукт %s добавлен. Цена добавлена %d\n", product, price)
}
