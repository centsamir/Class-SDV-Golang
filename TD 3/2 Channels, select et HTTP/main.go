package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//2.1
type Reponse struct {
	respText string
	err      error
}

func (a Reponse) String() string {
	response := a.respText
	if a.err != nil {
		response = a.err.Error()
	}
	return response
}

//2.3
func callServer(adr string, ch chan Reponse, wg *sync.WaitGroup) {
	defer wg.Done()
	//2.4
	res, err := http.Get(adr)
	//2.5
	if err != nil {
		erreur := Reponse{
			respText: "",
			err:      err,
		}
		ch <- erreur
		return
	} else if res.StatusCode != 200 {
		erreur := Reponse{
			respText: "",
			err:      errors.New("Le code retourné par le serveur indique une erreur: " + strconv.Itoa(res.StatusCode)),
		}
		ch <- erreur
	} else {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			erreur := Reponse{
				respText: "",
				err:      err,
			}
			ch <- erreur
		} else {
			reponse := Reponse{
				respText: string(body),
				err:      err,
			}
			ch <- reponse
		}

	}
	defer res.Body.Close()
}

func main() {
	//2.8
	var wg sync.WaitGroup
	//2.2
	ch1 := make(chan Reponse)
	ch2 := make(chan Reponse)
	wg.Add(2)
	//2.6- 2.7
	go callServer("http://localhost:4000/?id=id1", ch1, &wg)
	data := <-ch1
	fmt.Println("ch1: ", data)
	go callServer("http://localhost:4000/?id=id3", ch2, &wg)
	data = <-ch2
	fmt.Println("ch2: ", data)
	wg.Wait()
}
