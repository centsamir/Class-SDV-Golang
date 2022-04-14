package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var usersMap map[string]User

type User struct {
	Login    string `json:"userName"`
	Password string
	UserID   string `json:"userID"`
}

func (s *User) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	isUserID, userIDKey := searchId(r)
	sendResponse(w, isUserID, userIDKey)
}

func sendResponse(w http.ResponseWriter, isUserID bool, userIDKey string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if isUserID {
		w.WriteHeader(http.StatusOK)
		b1, _ := json.Marshal(usersMap[userIDKey])
		w.Write(b1)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func searchId(r *http.Request) (bool, string) {
	id := r.FormValue("id")
	isUserID := false
	userIDKey := ""
	for k := range usersMap {
		if k == id {
			isUserID = true
			userIDKey = k
			break
		}
	}
	return isUserID, userIDKey
}

func main() {
	file, _ := ioutil.ReadFile("users_end.json")

	var users []User
	json.Unmarshal(file, &users)

	fmt.Println(users)

	usersMap = make(map[string]User)
	for i := 0; i < len(users); i++ {
		usersMap[users[i].UserID] = users[i]
	}

	fmt.Println(usersMap)

	http.Handle("/", &User{})
	http.ListenAndServe("localhost:8080", nil)
}
