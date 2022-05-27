package main

import "fmt"

func main() {
	people1 := [10]string{"user1", "user2", "user3", "user4", "user5", "user6", "user7", "user8", "user9", "user10"}
	people2 := [10]string{"user11", "user12", "user13", "user14", "user15", "user16", "user17", "user18", "user19", "user20"}
	people3 := [10]string{"user21", "user22", "user23", "user24", "user25", "user26", "user27", "user28", "user29", "user30"}
	var peopleAll []string
	for _, value := range people1 {
		peopleAll = append(peopleAll, value)
	}
	for _, value := range people2 {
		peopleAll = append(peopleAll, value)
	}
	for _, value := range people3 {
		peopleAll = append(peopleAll, value)
	}

	for i := 0; i < len(peopleAll); i++ {
		fmt.Println(peopleAll[i])
	}
}
