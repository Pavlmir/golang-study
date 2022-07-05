package main

import "fmt"

//описываем новую структуру данных
type Car struct {
	Color  string
	Power  int
	Weight int
	Volume float32
	Mark   string
	Model  string
}

func main() {

	//создаем экземпляр структуры
	car := Car{Color: "Red", Power: 370, Weight: 1200, Volume: 4.3, Mark: "Cadillac", Model: "CTS"}
	car1 := Car{Color: "Blue", Power: 170, Weight: 1000, Volume: 2.4, Mark: "BMW", Model: "X1"}
	car2 := Car{Color: "Yellow", Power: 240, Weight: 1400, Volume: 3.3, Mark: "Toyota", Model: "Corola"}

	//создали список машин
	cars := []Car{}

	//заносим машины в список
	cars = append(cars, car)
	cars = append(cars, car1)
	cars = append(cars, car2)

	fmt.Println(cars)

	//вывести на экран список марок имеющихся авто
	for i := 0; i < len(cars); i++ {

		//выводим значение из ячейки Mark
		//fmt.Println(cars[i].Mark)
		if cars[i].Mark == "BMW" {

			fmt.Println("Нашли BMW")
		}

	}
}
