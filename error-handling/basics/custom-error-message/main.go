package main

import (
	"encoding/json"
	"errors"
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

	bs, err := toJSON(p1)

	if err != nil {
		// put corresponding to program logic error - check error basics
		log.Fatalln(err)
	}

	fmt.Println(string(bs))

}

// toJSON needs to return an custom error message also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)

	if err != nil {
		// return []byte{}, fmt.Errorf("Error at parsing json! (toJSON)\n", err) // creating custom error message
		return []byte{}, errors.New("error at parsing json! (toJSON)") // creating custom error 
	}

	return bs, err
}
