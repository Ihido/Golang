package main

import "fmt" //пакет функционала для форматирования и вывода данныхх в консоль.

func main() {
	fmt.Println()
}

var a, b, c string

func test() {
	var hello string
	hello = "My name"
	fmt.Println(hello)
}

func test1() {
	var (
		name string = "Mitya"
		age  int    = 23
	)
	fmt.Println(name)
	fmt.Println(age)
}

func integer() {
	var a int = 1
	fmt.Println(a)
}
