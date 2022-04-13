package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type User struct {
	Login    string `json:"userName"`
	Password string
}

func main() {
	user1 := User{Login: "Paul", Password: "pass123"}
	b1, _ := json.Marshal(user1)
	fmt.Println(string(b1))

	file, _ := ioutil.ReadFile("users.json")
	var users []User
	json.Unmarshal(file, &users)

	b0, _ := json.Marshal(users)
	fmt.Println(string(b0))

	for i := 0; i < len(users); i++ {
		b1, _ = json.Marshal(users[i])
		fmt.Println(string(b1))
	}
}
