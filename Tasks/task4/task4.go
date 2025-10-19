package main

import "fmt"

/*func main() {
	var rub float64
	fmt.Println("Введите сумму в рублях: ")
	fmt.Scanln(&rub)

	if rub <= 0 {
		fmt.Println("Сумма должна быть положительной!")
	}
	var usd = rub / 93.5
	var eur = rub / 101.2
	fmt.Printf("%.0f рублей = %.2f долларов, %2.f евро\n", rub, usd, eur)
}*/

/*func main() {
	var kg, m float64
	fmt.Println("Введите Ваш вес: ")
	fmt.Scanln(&kg)
	fmt.Println("Введите Ваш рост в метрах: ")
	fmt.Scanln(&m)
	if kg <= 0 {
		fmt.Println("Вес не может быть отрицательным!")
		return
	}
	if m <= 0 {
		fmt.Println("Рост не может быть отрицательным!")
		return
	}

	imt := kg / (m * m)

	var category string
	if imt < 18.5 {
		category = "недовес"
	} else if imt < 25 {
		category = "норма"
	} else if imt < 30 {
		category = "избыток"
	} else {
		category = "ожирение"
	}

	fmt.Printf("Ваш ИМТ: %.1f (%s)\n", imt, category)
}*/

func main() {
	var sum float64
	fmt.Println("Введите сумму покупки: ")
	fmt.Scanln(&sum)
	if sum <= 0 {
		fmt.Println("Сумма не может быть отрицательной!")
		return
	}
	var S float64
	if sum > 1000 {
		S = sum * 0.1
		sum = sum - S
		fmt.Printf("Итого: %.2f", sum)
	} else {
		fmt.Printf("Итого: %.2f", sum)
	}
}
