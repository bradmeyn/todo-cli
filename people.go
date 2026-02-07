package main

import (
	"errors"

	"github.com/google/uuid"
)


type Person struct {
    ID        string
    FirstName string
    LastName  string
    Email     string
    Age       int
}

func AddPerson(firstName, lastName, email string, age int) (Person, error) {

    if firstName == "" || lastName == "" || email == "" {
        return Person{}, errors.New("all fields are required")
    }

    id := uuid.New().String()

    person := Person{
        ID: id, 
        FirstName: firstName, 
        LastName: lastName, 
        Email: email, 
        Age: age,
    }

    return person, nil
}


