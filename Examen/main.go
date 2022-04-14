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

var Stask []Task

func main() {
	racine := func(rw http.ResponseWriter, _ *http.Request) {
		var ctask string
		// je crée ma string en parcourant tout les chants de ma slide avec le format souhaité
		for i, task := range Stask {
			if !task.Done {
				index := strconv.Itoa(i)
				ctask += "ID: " + index + ", task: " + "\"" + task.Description + "\"" + " \n"
			}
		}
		rw.WriteHeader(http.StatusOK)
		ctaskEnBytes := []byte(ctask)
		rw.Write(ctaskEnBytes)
	}
	done := func(rw http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			var ctask string
			// je crée ma string en parcourant tout les chants de ma slide avec le format souhaité
			for i, task := range Stask {
				if task.Done {
					index := strconv.Itoa(i)
					ctask += "ID: " + index + ", task: " + "\"" + task.Description + "\"" + " \n"
				}
			}
			rw.WriteHeader(http.StatusOK)
			ctaskEnBytes := []byte(ctask)
			rw.Write(ctaskEnBytes)
		case "POST":
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				fmt.Printf("Error reading body: %v", err)
				http.Error(rw, "can't read body", http.StatusBadRequest)
			}
			iBody, _ := strconv.Atoi(string(body))
			if iBody <= len(Stask) {
				Stask[iBody].Done = true
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
			sBody := string(body)
			task := Task{
				Description: sBody,
				Done:        false,
			}
			Stask = append(Stask, task)
			rw.WriteHeader(http.StatusOK)
		}
	}
	http.HandleFunc("/", racine)
	http.HandleFunc("/done", done)
	http.HandleFunc("/add", add)

	http.ListenAndServe("localhost:8080", nil)

}
