package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
)

var User3 map[string]User

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
func (user *User) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//7.a
	isUserID, userIDKey := searchId(r)
	//7.b
	sendResponse(w, isUserID, userIDKey)
}

func sendResponse(w http.ResponseWriter, isUserID bool, userIDKey string) {
	//7.c
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if isUserID {
		//7.d
		w.WriteHeader(http.StatusOK)
		//7.b
		b1, _ := json.Marshal(User3[userIDKey])
		//7.e
		w.Write(b1)
	} else {
		//7.d
		w.WriteHeader(http.StatusNotFound)
	}
}

//7.a
func searchId(r *http.Request) (bool, string) {
	id := r.FormValue("id")
	isUserID := false
	userIDKey := ""
	for k := range User3 {
		if k == id {
			isUserID = true
			userIDKey = k
			break
		}
	}
	return isUserID, userIDKey
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
	// User3 := make(map[string]string)
	// for i := 0; i < len(user2); i++ {
	// 	User3[user2[i].UserID] = user2[i].Login
	// }
	User3 = make(map[string]User)
	for i := 0; i < len(user2); i++ {
		User3[user2[i].UserID] = user2[i]
	}
	fmt.Println(User3)

	//2.5
	http.Handle("/", &User{})
	//2.6
	http.ListenAndServe("localhost:4000", nil)
	//2.7
}
