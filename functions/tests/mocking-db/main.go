package main

import (
	"fmt"
	"log"
)

// User represens user entry in db table
type User struct {
	ID int
	First string
}

// MockerDatastore is a temporary service that can stores retrievable data
type MockDatastore struct {
	Users map[int]User
}

// Methods for MockerDatastore struct 
func (md MockDatastore) GetUser(id int) (User, error) {
	user, ok := md.Users[id]
	if !ok {
		return User{}, fmt.Errorf("User %v not found", id)
	}
	return user, nil //nil is like null in other langs
}

func (md MockDatastore) SaveUser(u User) error {
	_, ok := md.Users[u.ID] // discard the value of this id - only check if it exists
	if ok {
		return fmt.Errorf("User %v already exists", u.ID)
	}
	md.Users[u.ID] = u
	return nil
}

// Datastore defines an interface for storing retrievable data.
// Any TYPE that implements this interface (has these two methods) is also of TYPE `Datastore`.
// This means any value of TYPE `MockDatastore` is also of TYPE `Datastore`
// This means we could have a value of TYPE `*sql.DB` and it can also be of TYPE `Datastore`
// This means we can write functions to take TYPE `Datastore` and use either of these:
// -- `MockDatastore`
// -- `*sql.DB`
// https://pkg.go.dev/database/sql#Open

type Datastore interface{
	GetUser(id int) (User, error)
	SaveUser(s User) error
}

// Service represents a service that stores retrievable data.
// We will attach methods to `Service` so that we can use either of these:
// -- `MockDatastore`
// -- `*sql.DB`
type Service struct {
	dataStore Datastore
}

// Methods for service struct
func (s Service) GetUser(id int) (User, error) {
	return s.dataStore.GetUser(id)
}

func (s Service) SaveUser(u User) error {
	return s.dataStore.SaveUser(u)
}

func main() {
	mocker_db := MockDatastore{
		Users: make(map[int]User),
	}

	db_service := Service{
		dataStore: mocker_db,
	}

	user1 := User{
		ID: 1,
		First: "Fifonż",
	}

	result := db_service.SaveUser(user1)

	if result != nil {
		log.Fatalf("Error %s", result)
	}

	user1Returend, err := db_service.GetUser(user1.ID)
	if err != nil {
		log.Fatalf("Error %s", result)
	}

	fmt.Println(user1)
	fmt.Println(user1Returend)
}