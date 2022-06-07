package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Должны быть экспортированные поля структур, с большой буквы
type Laptop struct {
	Name       string
	ScreenSize int
	Ssd        int
	Ram        int
}

type Phone struct {
	Name       string
	ScreenSize int
	Camera     int
	Ram        int
}

type Console struct {
	Name string
	Ram  int
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Получаем GET параметр
		var table string = r.URL.Query().Get("table")

		if table == "laptops" {
			// http://localhost:8080/?table=laptops
			laptop1 := Laptop{Name: "HONOR MagicBook", ScreenSize: 14, Ssd: 256, Ram: 16}
			laptop2 := Laptop{Name: "Lenovo", ScreenSize: 15, Ssd: 128, Ram: 32}
			laptop3 := Laptop{Name: "Asus", ScreenSize: 13, Ssd: 256, Ram: 16}
			laptops := []Laptop{laptop1, laptop2, laptop3}
			jsonData, _ := json.Marshal(laptops)
			fmt.Fprintf(w, string(jsonData))
		} else if table == "phones" {
			// http://localhost:8080/?table=phones
			phone1 := Phone{Name: "Samsung", ScreenSize: 6, Camera: 15, Ram: 16}
			phone2 := Phone{Name: "Xiaomi", ScreenSize: 5, Camera: 16, Ram: 32}
			phones := []Phone{phone1, phone2}
			jsonData, _ := json.Marshal(phones)
			fmt.Fprintf(w, string(jsonData))
		} else if table == "consoles" {
			// http://localhost:8080/?table=consoles
			console1 := Console{Name: "X-box", Ram: 64}
			console2 := Console{Name: "Sony", Ram: 32}
			consoles := []Console{console1, console2}
			jsonData, _ := json.Marshal(consoles)
			fmt.Fprintf(w, string(jsonData))
		} else {
			fmt.Fprintf(w, "Нужно выбрать категорию!")
		}
	})
	http.ListenAndServe(":8080", nil)
}
