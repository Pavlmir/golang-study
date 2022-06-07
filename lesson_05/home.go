/*
В Go использование заглавной буквы для любого объекта пакета (типа, переменной или функции) сделает его экспортированным автоматически.
export - Если имя поля начинается с заглавной буквы, оно будет доступно для чтения и записи с помощью кода вне пакета, где была определена структура.
local - Если имя поля начинается со строчной буквы, только код внутри пакета структуры сможет читать и записывать данные для этого поля.
*/

package main

import (
	"fmt"
	"strconv"
)

type Contacts struct {
	mobile uint64
	email  string
}

type SocialNetworks struct {
	name string
	link string
}

type User struct {
	id             uint64
	name           string
	surname        string
	age            uint8
	city           string
	university     string
	contacts       Contacts
	socialNetworks SocialNetworks
}

func main() {
	user1 := User{id: 1,
		name:       "Adel",
		surname:    "Yanis",
		age:        30,
		city:       "Moscow",
		university: "MSU",
		contacts: Contacts{mobile: 79001234545,
			email: "user1@test.com"},
		socialNetworks: SocialNetworks{name: "VK",
			link: "vk.com/user1"}}

	user2 := User{id: 2,
		name:       "Biber",
		surname:    "Jazov",
		age:        35,
		city:       "Irkuts",
		university: "ISU",
		contacts: Contacts{mobile: 79001236933,
			email: "user2@test.com"},
		socialNetworks: SocialNetworks{name: "FB",
			link: "fb.com/user2"}}

	user3 := User{id: 3,
		name:       "Cristian",
		surname:    "Ronald",
		age:        38,
		city:       "Sochi",
		university: "MTI",
		contacts: Contacts{mobile: 79001237788,
			email: "user3@test.com"},
		socialNetworks: SocialNetworks{name: "Instagram",
			link: "instagram.com/user3"}}

	users := []User{user1, user2, user3}
	for _, user := range users {
		fmt.Println(getInfoUser(user))
	}
}

func getInfoUser(user User) string {
	return fmt.Sprintln("Всем привет! Меня зовут " + user.name + " " + user.surname + ". Мне " + strconv.Itoa(int(user.age)) + " лет." +
		"Я живу в городе " + user.city + ". Я закончил " + user.university + ". Мой телефон " + strconv.Itoa(int(user.contacts.mobile)) + ", " +
		"e-mail: " + user.contacts.email + ". Я зарегистрирован в следующих соцсетях: " + user.socialNetworks.link + ".")
}
