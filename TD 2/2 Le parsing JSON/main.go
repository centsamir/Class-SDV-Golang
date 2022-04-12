package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Login    string `json:"userName"` //1.2
	Password string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//1.1.
	user := User{
		Login:    "Paul",
		Password: "pass123",
	}
	userJson, err := json.Marshal(user)
	check(err)
	//os.Stdout.Write(userJson)
	fmt.Printf("%s\n", userJson)

	//1.3
	data, err := os.ReadFile("./users.json")
	check(err)
	//fmt.Printf("%s\n", data)
	var user2 []User
	err = json.Unmarshal(data, &user2)
	check(err)
	fmt.Println(user2)
}
