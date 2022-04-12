package main

import (
	"fmt"
	"time"
)

type IPAddr [4]byte

//------------------//
//1.a
func (a IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", a[0], a[1], a[2], a[3])
}

//------------------//

//------------------//
//2.a
type MyError struct {
	When time.Time
	What string
}

//2.b
func (a MyError) Error() string {
	return fmt.Sprintf("%s : %s", a.When, a.What)
}

//2.c
func run() error {
	a := MyError{
		When: time.Now(),
		What: "Pomme",
	}
	return a
}

//------------------//

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
	//2.d

	err := run()
	if (MyError{} == err) {
		fmt.Println("It is an empty structure.")
	} else {
		fmt.Println(run())
	}
}
