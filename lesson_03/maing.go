package main

import "fmt"

func main() {

	//создаем массив  (срез)
	var numbers []int
	fmt.Println(numbers)

	//добавляем элемент в конец слайса
	//append("куда добвляем", "что добавляем") - добавляет в конец среза
	numbers = append(numbers, 25)
	numbers = append(numbers, 25)
	numbers = append(numbers, 25)
	numbers = append(numbers, 25)
	numbers = append(numbers, 25)
	numbers = append(numbers, 15)
	numbers = append(numbers, 17)
	numbers = append(numbers, 17)
	numbers = append(numbers, 17)
	numbers = append(numbers, 17)
	numbers = append(numbers, 17)
	numbers = append(numbers, 5)
	numbers = append(numbers, 42)
	numbers = append(numbers, 42)

	numbers = append(numbers, 42)

	fmt.Println(numbers)
	//считаем сумму
	fmt.Println(getSum(numbers))

	//len(маccив/срез) - подсчет количества элементов массива/среза
	fmt.Println(len(numbers))

	//выводим среднее
	fmt.Println(avarge(numbers))

	/*
		//наполняем данными
		numbers[0] = 25
		numbers[1] = 15
		numbers[2] = 17
		numbers[3] = 5
		numbers[4] = 42
		fmt.Println(numbers)

		//выводим сумму массива
		fmt.Println(getSum(numbers))
	*/
}

func getSum(array []int) int {

	//return array[0] + array[1] + array[2] + array[3] + array[4]
	var sum int = 0

	//бежим по всем ячейкам среза
	//for СТАРТ; ФИНИШ; ШАГ
	//for 0; 0 < 6; 0 + 1
	//for 1; 1 < 6; 1 + 1
	//for 2; 2 < 6; 2 + 1
	//for 3; 3 < 6; 3 + 1
	//for 4; 4 < 6; 4 + 1
	//for 5; 5 < 6; 5 + 1
	for i := 0; i < len(array); i = i + 1 {

		//смотрю текущую ячейку и прибавляю к общей сумме
		sum = sum + array[i]
	}

	return sum
}

func avarge(array []int) int {

	return getSum(array) / len(array)
}
