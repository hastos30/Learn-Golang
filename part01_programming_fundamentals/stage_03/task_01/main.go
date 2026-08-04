package main

import "fmt"

func main() {
	evenCount := 0
	oddSum := 0

	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			oddSum += i
			fmt.Println("Нечётное:", i)
		} else {
			evenCount++
			fmt.Println("Чётное:", i)
		}
	}

	fmt.Println("Количество чётных:", evenCount)
	fmt.Println("Сумма нечётных:", oddSum)
}
