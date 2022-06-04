package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Laptop struct {
	Mark  string
	Model string
}

type Phone struct {
	Mark  string
	Model string
}

type Console struct {
	Mark  string
	Model string
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		//создаем 3 объекта
		laptop := Laptop{Mark: "DELL", Model: "G5"}
		phone := Phone{Mark: "Xiaomi", Model: "Mi Max 3"}
		console := Console{Mark: "Sony", Model: "Playstation 5"}

		//получаем GET параметр
		table := r.URL.Query().Get("table")

		//выводим данные в зависимости от знаения GET параметра\
		if table == "laptops" {

			//преобразуем обьект в JSON
			jsonData, _ := json.Marshal(laptop)
			//выводим на экран
			fmt.Fprintf(w, string(jsonData))

		} else if table == "phones" {

			//преобразуем обьект в JSON
			jsonData, _ := json.Marshal(phone)
			//выводим на экран
			fmt.Fprintf(w, string(jsonData))

		} else if table == "consoles" {

			//преобразуем обьект в JSON
			jsonData, _ := json.Marshal(console)
			//выводим на экран
			fmt.Fprintf(w, string(jsonData))

		} else {
			fmt.Fprintf(w, "Нужно ввести параметр")
		}

	})
	http.ListenAndServe(":8080", nil)
}
