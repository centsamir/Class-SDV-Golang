package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

type User struct {
	Login    string `json:"userName"` //1.2
	Password string
	UserID   string `json:"userID"` //2.2
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//------------------//
//2.4
func ServeHTTP(writer http.ResponseWriter, request *http.Request) {

}

//------------------//

func main() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
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
	User3 := make(map[string]string)
	for i := 0; i < len(user2); i++ {
		User3[user2[i].UserID] = user2[i].Login
	}
	fmt.Println(User3)

	//2.4
	http.HandleFunc("/", ServeHTTP())
}
