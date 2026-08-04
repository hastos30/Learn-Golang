package main

import "fmt"

func main() {
	/*
		Часть 1
	*/

	// var name string = "Виктор"
	// age := 31

	// name = "Viktor"
	// age = 32

	// fmt.Println(name)
	// fmt.Println(age)

	/*
		Часть 2
	*/

	// var age int = 31
	// // var temperature float64
	// temperature := 24.5
	// // var isLearning bool
	// isLearning := true
	// // var language string
	// language := "Go"

	// fmt.Printf("%T\n", age)
	// fmt.Printf("%T\n", temperature)
	// fmt.Printf("%T\n", isLearning)
	// fmt.Printf("%T\n", language)

	/*
		Часть 3
	*/

	// const hoursPerDay = 24

	// Ошибка (не можем изменить константу)
	// hoursPerDay = 25

	// days := 7
	// totalHours := days * hoursPerDay

	// fmt.Println("Дней:", days)
	// fmt.Println("Часов:", totalHours)

	/*
		Часть 4
	*/

	// a := 10
	// b := 3

	// a := 10.0
	// b := 3.0

	// fmt.Println("Сложение:", a+b)
	// fmt.Println("Вычитание:", a-b)
	// fmt.Println("Умножение:", a*b)
	// fmt.Println("Деление:", a/b)
	// fmt.Println("Остаток:", a%b)

	/*
		Часть 5
	*/

	// total := 10
	// count := 4

	// average := float64(total) / float64(count)

	// fmt.Println("Среднее:", average)
	// fmt.Printf("Тип total: %T\n", total)
	// fmt.Printf("Тип average: %T\n", average)

	/*
		Часть 6
	*/

	// age := 20
	// minimumAge := 18

	// fmt.Println(age > minimumAge)
	// fmt.Println(age < minimumAge)
	// fmt.Println(age == minimumAge)
	// fmt.Println(age != minimumAge)
	// fmt.Println(age >= minimumAge)
	// fmt.Println(age <= minimumAge)

	/*
		Часть 7
	*/

	// age := 20
	// hasPermission := false

	// // canEnter := age >= 18 && hasPermission
	// canEnter := age >= 18 || hasPermission

	// fmt.Println(canEnter)

	// number := 10
	// divisor := 0

	// canDivide := divisor != 0 && number/divisor > 2

	// fmt.Println(canDivide)

	// isBlocked := false
	// canEnter := !isBlocked

	// fmt.Println(canEnter)

	const minimumAge = 18

	age := 20
	hasTicket := true
	isBlocked := false

	canEnter := age >= minimumAge && hasTicket && !isBlocked

	fmt.Println("Возраст:", age)
	fmt.Println("Есть билет:", hasTicket)
	fmt.Println("Заблокирован:", isBlocked)
	fmt.Println("Можно войти", canEnter)
}
