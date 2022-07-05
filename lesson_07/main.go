package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		//получаем GET параметр
		var login string = r.URL.Query().Get("login") // в PHP - $_GET['name']
		var password string = r.URL.Query().Get("password")

		if login == "inordic" && password == "123" {
			fmt.Fprintf(w, "Поздравляем вы успешно авторизовались")
		} else {
			fmt.Fprintf(w, "Неверный логин или пароль")
		}

		//выводим на экран

		//fmt.Fprintf(w, "<a href='https://yandex.ru' ><img src='https://akibamarket.com/wp-content/uploads/2022/05/104303-una-chica-de-anime-se-vuelve-real-en-una-fabulosa-edicion-de-video.png'></a>")
		//fmt.Fprintf(w, "<img src='https://akibamarket.com/wp-content/uploads/2022/05/104303-una-chica-de-anime-se-vuelve-real-en-una-fabulosa-edicion-de-video.png'>")
	})
	http.ListenAndServe(":8080", nil)
}
