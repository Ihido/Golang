package main

import "fmt"

//-----------------------------------------------------1

// func main() {
// 	var name, city string
// 	var age int
// 	fmt.Println("Введите Ваше имя: ")
// 	fmt.Scanln(&name)
// 	fmt.Println("Введите Ваш возраст: ")
// 	fmt.Scanln(&age)
// 	fmt.Println("Введите Ваш город: ")
// 	fmt.Scanln(&city)
// 	fmt.Println("Имя:", name, "Возраст:", age, "Город:", city)
// }

//-----------------------------------------------------2

// func main() {
// 	var word string
// 	var count int
// 	fmt.Println("Введите слово: ")
// 	fmt.Scanln(&word)
// 	fmt.Println("Введите количество повторений: ")
// 	fmt.Scanln(&count)
// 	enter := ""
// 	for i := 0; i < count; i++ {
// 		enter += word
// 	}

// 	fmt.Println(enter)
// }

//-----------------------------------------------------3

func main() {
	var n int
	fmt.Println("Введите число:")
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(i)
		}
		fmt.Println()
	}
}
