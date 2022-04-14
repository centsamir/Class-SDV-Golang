package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Task struct {
	Description string
	Done        bool
}

var Tasks []Task

func main() {
	racine := func(rw http.ResponseWriter, _ *http.Request) {
		var stringTasks string
		// je crée ma string en parcourant tout les chants de ma slide avec le format souhaité
		for i, task := range Tasks {
			if !task.Done {
				index := strconv.Itoa(i)
				stringTasks += "ID: " + index + ", task: " + "\"" + task.Description + "\"" + " \n"
			}
		}
		rw.WriteHeader(http.StatusOK)
		tasksEnBytes := []byte(stringTasks)
		rw.Write(tasksEnBytes)
	}
	done := func(rw http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			var stringTasks string
			// je crée ma string en parcourant tout les chants de ma slide avec le format souhaité
			for i, task := range Tasks {
				if task.Done {
					index := strconv.Itoa(i)
					stringTasks += "ID: " + index + ", task: " + "\"" + task.Description + "\"" + " \n"
				}
			}
			rw.WriteHeader(http.StatusOK)
			tasksEnBytes := []byte(stringTasks)
			rw.Write(tasksEnBytes)
		case "POST":
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Printf("Error reading body: %v", err)
				http.Error(rw, "can't read body", http.StatusBadRequest)
			}
			id, _ := strconv.Atoi(string(body))
			if id <= len(Tasks) {
				Tasks[id].Done = true
			} else {
				rw.WriteHeader(http.StatusBadRequest)
			}
		default:
			rw.WriteHeader(http.StatusBadRequest)
		}
	}
	add := func(rw http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			rw.WriteHeader(http.StatusBadRequest)
		} else {
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Printf("Error reading body: %v", err)
				http.Error(rw, "can't read body", http.StatusBadRequest)
			}
			description := string(body)
			task := Task{
				Description: description,
				Done:        false,
			}
			Tasks = append(Tasks, task)
			rw.WriteHeader(http.StatusOK)
		}
	}
	http.HandleFunc("/", racine)
	http.HandleFunc("/done", done)
	http.HandleFunc("/add", add)

	http.ListenAndServe("localhost:8080", nil)

}
