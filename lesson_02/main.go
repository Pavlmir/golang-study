package main

import (
	"fmt"
)

//подключаем пакет для вывода данных

//главное действие - есть всегда
func main() {

	var first int = 10
	var second int = 2

	//выводим сумму двух чисел
	fmt.Println(add(first, second))

	//выводим разность двух чисел

	//выводим произведение чисел

	//выводим деление чисел

	//вывод сразу 4 действий
	fmt.Println(math(first, second))

	fmt.Println(generateText(3, 5, 7, 1))
}

//создаю функцию сложения
func add(a int, b int) int {

	//производим ДЕЙСТВИЕ с ИСХОДНЫМИ ДАННЫМИ
	var result int = a + b

	//выдаем РЕЗУЛЬТАТ работы
	return result
}

func math(a int, b int) (int, int, int, int) {

	//делаем 4 разных коробки и наполняем результатом
	var result1 int = a + b
	var result2 int = a - b
	var result3 int = a * b
	var result4 int = a / b

	//выдаем все 4 коробки сразу
	return result1, result2, result3, result4

}
