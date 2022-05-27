package main

import "fmt"

func main() {

	//срез с именами
	var names []string = []string{"Анна", "Настя", "Олег", "Анна"}

	//переменная для поиска имени
	var searchName string = "Джон"

	//выясняем есть ли в списке Олег
	search(names, searchName)

	//считаем сколько у нас Анн
	fmt.Println("Найдно "+searchName+" : ", searchCount(names, searchName))
}

//функция для поиска по имени
//names - там ГДЕ ищем
// target - кого ищем
func search(names []string, target string) bool {

	//пробежаться в цикле по всем элементам массива(среза)
	for i := 0; i < len(names); i++ {

		//проверяем
		if names[i] == target {

			fmt.Println("Нашли совпадение")
			return true
		}
	}

	//если всех пербрали и не нашли
	return false

}

//функция для подсчета людей с определенным именем
func searchCount(names []string, target string) int {

	//создаем счетчик имен
	var count int = 0

	//просмотреть все имена
	for i := 0; i < len(names); i++ {

		//если нашли имя
		if names[i] == target {

			//то увеличиваем счетчик на 1
			count = count + 1
		}
	}

	//возвращаем сколько всего нашли человек с таким именем
	return count
}
