package main

//подключаем новый пакет
import "fmt"

func main() {

	//вывод данных в консоль
	//Println() - вывод в консоль
	fmt.Println("Привет Nordic School")

	//создаем переменную
	//var first int = 9876
	first := 9876

	//var second int = 3245
	second := 3245

	//var result int = first * second
	result := first * second

	fmt.Println(result)

	name := "Иван"
	lastName := "Жевакин"

	fullName := name + " " + lastName
	fmt.Println(fullName)
}
