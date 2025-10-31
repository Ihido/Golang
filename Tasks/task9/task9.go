package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
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

// Задачка далась не сра зу. Для решения необходимо разбить по этапам.
// 1. Определиться с типом возвращаемого значения: в данном случае bool
// 2. Каким образом мы будем проверять длину строки? спомощью встроенной фнкуции len
// 3. Как мы будем проверять строку на наличие цифр? Спомощью цикла фор, предварительно создав "флажок number", который будет говорить "Есть ли цифра?"
// func ValidPassword(password string) bool {
// 	if len(password) < 8 {
// 		return false
// 	}
// 	number := false
// 	for i := 0; i < len(password); i++ {
// 		if password[i] >= '0' && password[i] <= '9' {
// 			number = true
// 			break
// 		}
// 	}
// 	return number
// }

// func main() {
// 	result := ValidPassword("Grinf8")
// 	fmt.Print(result)
// }

//-----------------------------------------------------2.2

// func SumDigits(number int) int {
// 	sum := 0 // Создали флажок для записи чисел

// 	for number > 0 {
// 		digit := number % 10 // Отбираем последнюю цифру (123 = 12 % 10 = 3)
// 		sum += digit         // Записываем последнее число в переменную
// 		number = number / 10 // Цикл продолжается со следующей цифры 12
// 	}

// 	return sum
// }

// func main() {
// 	result1 := SumDigits(123)
// 	fmt.Println(result1) // 6

// 	result2 := SumDigits(987)
// 	fmt.Println(result2) // 24 (9+8+7)

// 	result3 := SumDigits(1000)
// 	fmt.Println(result3) // 1
// }

//-----------------------------------------------------2.3

// func IsPrime(number int) bool {
// 	if number < 2 {
// 		return false
// 	}

// 	for i := 2; i*i == number; i++ {
// 		if number%i == 0 {
// 			return false
// 		}

// 	}
// 	return true
// }
// func main() {
// 	message1 := IsPrime(7)
// 	fmt.Println(message1)

// 	message2 := IsPrime(8)
// 	fmt.Println(message2)
// }

//-----------------------------------------------------2.4

// func GenerateEmail(name, domain string) string { //нельзя использовать print, так как результат будет false(потому что сравнивает пустую строку email с name+domain)
// 	email := name + "@" + domain
// 	return email
// }

// func main() {
// 	message1 := GenerateEmail("Иван", "mail.ru")
// 	message2 := GenerateEmail("Дмитрий", "yandex.ru")
// 	message3 := GenerateEmail("Олег", "gmail.com")
// 	fmt.Println(message1)
// 	fmt.Println(message2)
// 	fmt.Println(message3)
// }

//-----------------------------------------------------2.5

func TextStats(text string) (int, int) {
	charCount := utf8.RuneCountInString(text)

	words := strings.Fields(text)
	warCount := len(words)

	return warCount, charCount
}

func main() {
	words, chars := TextStats("Salut world")
	fmt.Printf("%d слова, %d символов\n", words, chars)

	words2, chars2 := TextStats("Привет мир кекВ!")
	fmt.Printf("Слова: %d"+" "+"Символов: %d", words2, chars2)
}
