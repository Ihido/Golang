package main

import (
	"fmt"
)

//-----------------------------------------------------1.1

// func Square(num int) int {
// 	return num * num
// }

// func main() {
// 	result := Square(5)
// 	fmt.Print(result)
// }

//-----------------------------------------------------1.2

// func ConvertUSDToRUB(usd int) int {
// 	return usd * 90
// }

// func main() {
// 	result := ConvertUSDToRUB(10)
// 	fmt.Print(result)
// }

//-----------------------------------------------------1.3

// func TimeGreeting(hour int) string {
// 	if hour <= 11 {
// 		return "Доброе утро"
// 	}
// 	if hour <= 18 {
// 		return "Добрый день"
// 	} else {
// 		return "Добрый вечер"
// 	}
// }

// func main() {
// 	message := TimeGreeting(14)
// 	fmt.Print(message)
// }

//-----------------------------------------------------2.1

// Задачка далась не сразу. Для решения необходимо разбить по этапам.
// 1. Определиться с типом возвращаемого значения: в данном случае bool
// 2. Каким образом мы будем проверять длину строки? спомощью встроенной фнкуции len
// 3. Как мы будем проверять строку на наличие цифр? Спомощью цикла фор, предварительно создав "флажок number", который будет говорить "Есть ли цифра?"
func ValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}
	number := false
	for i := 0; i < len(password); i++ {
		if password[i] >= '0' && password[i] <= '9' {
			number = true
			break
		}
	}
	return number
}

func main() {
	result := ValidPassword("Grinf8")
	fmt.Print(result)
}
