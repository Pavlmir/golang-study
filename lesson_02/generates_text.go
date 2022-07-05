package main

import "fmt"

func main() {
	generates_text("Audi", 200, 100, "красный")
	generates_text("BMW", 180, 90, "желтый")
	generates_text("Kia", 150, 80, "серый")
	generates_text("Fiat", 300, 110, "малиновый")
	generates_text("Bentley", 250, 75, "розовый")
}

func generates_text(brand string, power int, volume int, color string) {
	fmt.Printf("Это отличный автомобиль марки %s. "+
		"Мощность его двигателя составляет %d лошадиных сил. "+
		"У нее большой двигатель объемом %d л. "+
		"Прекрасный %s цвет - подчеркивает ее элегантность\n",
		brand, power, volume, color)
}
