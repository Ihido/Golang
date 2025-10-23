package main

import "fmt"

//-----------------------------------------------------1

// func Greet(name string) string {
// 	return "Привет, " + name + "!"
// }

// func main() {
// 	message1 := Greet("Анна")
// 	fmt.Println(message1)

// 	message2 := Greet("Анна")
// 	fmt.Println(message2)

// 	fmt.Println(Greet("Мария"))
// }

//-----------------------------------------------------2

// func Add(num1, num2 int) int {
// 	return num1 + num2
// }

// func main() {
// 	result1 := Add(5, 3)
// 	fmt.Println(result1)
// }

//-----------------------------------------------------3

// func IsEven(num int) bool {
// 	if num%2 == 0 {
// 		return true
// 	} else {
// 		return false
// 	}
// }

// func main() {
// 	result1 := IsEven(4)
// 	fmt.Println(result1)

// 	result2 := IsEven(7)
// 	fmt.Println(result2)
// }

//-----------------------------------------------------4

// func Max(num1, num2 int) int {
// 	if num1 > num2 {
// 		return num1
// 	} else {
// 		return num2
// 	}
// }

// func main() {
// 	result := Max(3, 8)
// 	fmt.Print(result)
// }

//-----------------------------------------------------5

func Repeat(text string, count int) string {
	result := ""
	for i := 1; i <= count; i++ {
		result += text
	}
	return result
}

func main() {
	result := Repeat("Go", 3)
	fmt.Print(result)
}
