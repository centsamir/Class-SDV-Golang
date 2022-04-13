package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Login    string `json:"userName"` //1.2
	Password string
	UserID   int `json:"userID"` //2.2
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//1.1.
	fmt.Println("1.1")
	user := User{
		Login:    "Paul",
		Password: "pass123",
	}
	userJson, err := json.Marshal(user)
	check(err)
	//os.Stdout.Write(userJson)
	fmt.Printf("%s\n", userJson)

	//1.3
	fmt.Println("1.3")
	data, err := os.ReadFile("./users.json")
	check(err)
	fmt.Printf("%s\n", data)
	var user2 []User
	err = json.Unmarshal(data, &user2)
	check(err)
	fmt.Println(user2)

	//2.3
	fmt.Println("2.3")
	var User3 map[string]User
	fmt.Println(User3)
	fmt.Println(user2[0])
	for _, elem := range user2 {
		User3["Login"] = 

	}
	
}
