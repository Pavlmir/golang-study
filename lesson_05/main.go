package main

import "fmt"

type Laptop struct {
	Mark       string
	Model      string
	RAM        Ram
	CPU        Cpu
	GPU        Gpu
	Weight     float32
	Resolution string
	Wifi       []string
}

type Ram struct {
	Frequency uint16
	Volume    uint8
}

type Cpu struct {
	Core      byte
	Frequency float32
}

type Gpu struct {
	Frequency float32
	Volume    byte
}

func main() {

	laptop1 := Laptop{Mark: "Dell", Model: "G5", RAM: Ram{Frequency: 2500, Volume: 16}, CPU: Cpu{Core: 8, Frequency: 2.4}, GPU: Gpu{Frequency: 2700, Volume: 6}, Wifi: []string{"802.11n", "802.11f"}}

	fmt.Println(laptop1)

	fmt.Println(getDescription(laptop1))
}

func getDescription(laptop Laptop) string {

	var text string = "Частота процессора " + fmt.Sprint(laptop.CPU.Frequency)
	return text
}
