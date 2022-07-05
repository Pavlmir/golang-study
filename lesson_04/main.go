package main

//подлючаем пакет fmt чтобы выводить в консоль
import "fmt"

func main() {

	var namesAll [10]string
	namesAll[0] = "Николай"
	namesAll[1] = "Джон"
	namesAll[2] = "Артем"
	namesAll[3] = "Кевин"
	namesAll[4] = "Андрей"
	namesAll[5] = "Марина"
	namesAll[6] = "Анна"
	namesAll[7] = "Катя"
	namesAll[8] = "Софья"
	namesAll[9] = "Диана"

	//коробка под отсортированные
	var namesSelected []string

	//начинаем преебирать все имена
	for i := 0; i < len(namesAll); i++ {

		//включаем фильтр
		// != НЕ РАВНО
		if namesAll[i] != "Николай" {

			//добаляем имя в коробку с отсортированными
			namesSelected = append(namesSelected, namesAll[i])
		}
	}

	//выводим отсортированные имена на экран
	fmt.Println(namesSelected)

}
