package main

import (
	"fmt"
	"time"
)

// https://go.dev/tour/methods/16
// https://go.dev/tour/methods/19

type IPAddr [4]byte

type MyError struct {
	When time.Time
	What string
}

func (er *MyError) String() string {
	return "rpeoir"
}

func PrintIt(input interface{}) {
	fmt.Println(input)
}

func PrintItType(input interface{}) {
	switch v := input.(type) {
	case int:
		fmt.Printf("%v is an int\n", v)
	case string:
		fmt.Printf("%q is a string of %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{time.Now(), "it didn't work"}
}

// Sprintf
func (ip IPAddr) String() string {
	a := make([]interface{}, len(ip))
	for i, x := range ip {
		a[i] = x
	}
	return fmt.Sprintf("%d.%d.%d.%d", a...)
}

// Strconv
// func (ip IPAddr) String() string {
// 	s := make([]string, len(ip))
// 	for i, val := range ip {
// 		s[i] = strconv.Itoa(int(val))
// 	}
// 	return strings.Join(s, ".")
// }

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	if err := run(); err != nil {
		fmt.Println(err)
	}

	PrintIt(5)
	PrintIt("Ceci est une chaîne de caractères")
	PrintItType(4)
	PrintItType("Ceci est une chaîne de caractères")
}
