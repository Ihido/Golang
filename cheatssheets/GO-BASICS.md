# 🎯 Основы Go для начинающих

## 🏗️ Структура программы
```go
package main

import "fmt"

func main() {
    fmt.Println("Привет, мир!")
}
```

## 📦 Переменные и типы данных
```go
// Объявление переменных
var name string = "Алексей"
var age int = 25
var isProgrammer bool = true

// Короткое объявление (чаще используется)
name := "Алексей"
age := 25
price := 19.99
```

## 🔤 Основные типы данных
```go
var text string = "Текст"       // Строка
var number int = 42             // Целое число
var decimal float64 = 3.14      // Дробное число
var isTrue bool = true          // Логический тип
```

## 🔄 Функции
```go
// Простая функция
func sayHello() {
    fmt.Println("Привет!")
}

// Функция с параметрами
func greet(name string) {
    fmt.Println("Привет,", name)
}

// Функция с возвращаемым значением
func add(a int, b int) int {
    return a + b
}
```

## 🎮 Условия (if/else)
```go
age := 18

if age >= 18 {
    fmt.Println("Ты взрослый")
} else {
    fmt.Println("Ты еще молод")
}

// Короткая запись
if temperature := 25; temperature > 30 {
    fmt.Println("Жарко!")
}
```

## 🔁 Циклы (for)
```go
// Обычный цикл
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// Бесконечный цикл
for {
    fmt.Println("Бесконечность!")
    break // чтобы выйти
}

// Цикл по массиву
numbers := []int{1, 2, 3}
for index, value := range numbers {
    fmt.Println(index, value)
}
```

## 📚 Массивы и срезы (slices)
```go
// Массив (фиксированный размер)
var numbers [3]int = [3]int{1, 2, 3}

// Срез (динамический массив)
names := []string{"Аня", "Боря", "Вова"}
names = append(names, "Глаша")  // Добавить элемент

// Работа со срезами
fmt.Println(names[0])           // Первый элемент
fmt.Println(names[1:3])         // Срез со 2 по 3 элемент
```

## 🗂️ Структуры (struct)
```go
// Определение структуры
type Person struct {
    Name string
    Age  int
}

// Создание экземпляра
person := Person{
    Name: "Мария", 
    Age:  30,
}

// Доступ к полям
fmt.Println(person.Name)
```

## 🐞 Частые ошибки новичков
```go
// ❌ Неиспользуемые переменные (ошибка компиляции)
var unused int = 10

// ✅ Правильно - используй все переменные
used := 10
fmt.Println(used)

// ❌ Неправильные типы
var count int = "текст"  // Ошибка!

// ✅ Правильные типы
var count int = 42
var text string = "текст"
```

## 💡 Полезные советы
- Используй `go fmt` для форматирования кода
- Имена переменных должны быть понятными (`userName`, а не `un`)
- Комментируй сложные части кода
- Тестируй каждую функцию отдельно