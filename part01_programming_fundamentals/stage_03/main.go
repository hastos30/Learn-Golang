package main

import "fmt"

func main() {
	// age := 21
	// isVerified := false

	// if age >= 18 && isVerified {
	// 	fmt.Println("Доступ разрешён")
	// } else {
	// 	fmt.Println("Доступ запрещён")
	// }

	// isAdmin := false
	// isManager := true

	// if isAdmin || isManager {
	// 	fmt.Println("Доступ к панели разрешен")
	// } else {
	// 	fmt.Println("Доступ запрещен")
	// }

	// isBlocked := false

	// if !isBlocked {
	// 	fmt.Println("Пользователь может войти")
	// }

	// day := 3

	// switch day {
	// case 1:
	// 	fmt.Println("Понедельник")
	// case 2:
	// 	fmt.Println("Вторник")
	// case 3:
	// 	fmt.Println("Среда")
	// default:
	// 	fmt.Println("Неизвестный день")
	// }

	// day := 10

	// switch day {
	// case 1, 2, 3, 4, 5:
	// 	fmt.Println("Рабочий день")
	// case 6, 7:
	// 	fmt.Println("Выходной")
	// default:
	// 	fmt.Println("Некорректный номер дня")
	// }

	// status := "paid"

	// switch status {
	// case "new":
	// 	fmt.Println("Заказ создан")
	// case "paid":
	// 	fmt.Println("Заказ оплачен")
	// case "cancelled":
	// 	fmt.Println("Заказ отменён")
	// default:
	// 	fmt.Println("Неизвестный статус")
	// }

	// for i := 1; i <= 5; i++ {
	// 	fmt.Println(i)
	// }

	// for i := 0; i <= 10; i += 2 {
	// 	fmt.Println(i)
	// }

	// count := 5

	// for count >= 1 {
	// 	fmt.Println(count)
	// 	count--
	// }

	// count := 1

	// for {
	// 	fmt.Println(count)

	// 	if count == 3 {
	// 		break
	// 	}

	// 	count++
	// }

	// for i := 1; i <= 10; i++ {
	// 	if i%2 != 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	// count := 0
	// sum := 0

	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		sum += i
	// 		count++
	// 	}
	// }

	// fmt.Println("Количество:", count)
	// fmt.Println("Сумма:", sum)

	// for row := 1; row <= 3; row++ {
	// 	for column := 1; column <= 2; column++ {
	// 		fmt.Println("Строка:", row, "Столбец:", column)
	// 	}
	// }

	// for row := 1; row <= 4; row++ {
	// 	for column := 1; column <= 7; column++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	// for i := 1; i <= 5; i++ {
	// 	for j := 1; j <= i; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	// for i := 5; i >= 1; i-- {
	// 	for j := 0; j < i; j++ {
	// 		fmt.Print("*")

	// 	}
	// 	fmt.Println()
	// }

	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
}
