package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
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

//2.3
func callServer(adr string, ch chan Reponse) {
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
	//2.2
	ch1 := make(chan Reponse)
	ch2 := make(chan Reponse)

}
