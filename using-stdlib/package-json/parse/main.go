package main

import (
	"encoding/json"
	"fmt"
)

// struct fields need to be exported to marshall them correctly
type user struct {
	First string
	Age   int
}

// Always check documentation on sertain stdlib functions
func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here

	json_users, err := json.Marshal(users)

	if err != nil {
		fmt.Println("Parsing went wrong!")
	}
	fmt.Println(string(json_users)) // parse it to string - Marshal returns byte array ([]byte)
}
