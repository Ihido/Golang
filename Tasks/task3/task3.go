package main

import "fmt"

func main() {
	var num int
	fmt.Println("Введите число: ")
	fmt.Scanln(&num)
	if num%2 == 0 {
		fmt.Print("Число ", num, " четное!")
	} else {
		fmt.Print("Число ", num, " не чётное!")
	}
}
