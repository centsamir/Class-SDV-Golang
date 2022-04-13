package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go maFonction(&wg)
	fmt.Println("Fin du programme")
	wg.Wait()
}

func maFonction(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("j'ai fini !")
}
