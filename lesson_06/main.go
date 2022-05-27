package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name     string
	LastName string
	Age      int
}

func main() {

	//активируем веб-сервер
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		//создаем юзера
		user1 := User{Name: "John", LastName: "Smith", Age: 23}

		//скармливаем ОБЪЕКТ(структура) функции которая преобраует его в json СТРОКА(string)
		jsonData, _ := json.Marshal(user1)

		//выводим в консоль
		//string() - функция для преобразования СРЕЗА байт в строку
		fmt.Println(string(jsonData))

		//выводим в ОТВЕТ
		//Fprintf() - вывод в ОТВЕТ
		fmt.Fprintf(w, string(jsonData))
	})
	http.ListenAndServe(":8080", nil)

}
