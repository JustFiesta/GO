package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	// bs, _ := json.Marshal(p1) // handle the possible error

	bs, err := json.Marshal(p1) // stringify json data

	// in this program handling the error is simply close the program if marshalling does not work
	if err != nil {
		log.Println(err)
		// fmt.Println(err) output error to std output - can recover from
		// log.SetOutput() write log to file - can recover from
		// log.Panic() print to std and panic  - can recover from
		// panic() - can recover from
		// log.Fatalln() print to std and sys.exit(1) - function caput, no recovery from this

		return
	}

	fmt.Println(string(bs))

}
