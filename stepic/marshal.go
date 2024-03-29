package main

import (
	"fmt"
	"time"
	"encoding/json"
	"log"
)

type Person struct {
	Name        string    `json:"Имя"`
	Email       string    `json:"Почта"`
	DateOfBirth time.Time `json:"-"` // - означает, что это поле не будет сериализовано
	}
	
	func main() {
	
		man := Person{
			Name:        "Alex",
			Email:       "alex@yandex.ru",
			DateOfBirth: time.Now(),
		}
		jsMan, err := json.Marshal(man)
		if err != nil {
			log.Fatalln("unable marshal to json")
		}
		fmt.Printf("Man %v", string(jsMan)) // Man {"Имя":"Alex","Почта":"alex@yandex.ru"}
	} 