package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/google/uuid"
)


type Person struct {
    ID        string
    FirstName string
    LastName  string
    Email     string
    Age       int
}

func AddPerson(firstName, lastName, email string, age int) error {

    if firstName == "" || lastName == "" || email == "" {
        return  errors.New("all fields are required")
    }

    id := uuid.New().String()

    person := Person{
        ID: id, 
        FirstName: firstName, 
        LastName: lastName, 
        Email: email, 
        Age: age,
    }

    // Convert person struct to a slice of strings for CSV storage
    record := []string{person.ID, person.FirstName, person.LastName, person.Email, strconv.Itoa(person.Age) }

    err := AppendCSV("db.csv", record)
    if err != nil {
        return  fmt.Errorf("failed to append person to CSV: %w", err)
    }

    return  nil
}

func DeletePerson(id string) error {

    records, err := ReadCSV("db.csv")
    if err != nil {
        fmt.Errorf("Failed to read CSV")
    }

    var updated_records [][]string
    for _, record := range records {

        if record[0] == id {
            fmt.Println("Removing /%s /%s", record[1], record[2])
        } else {
            updated_records = append(updated_records, record)
        }
    }

    err = WriteCSV("db.csv", updated_records)
        if err != nil {
        return fmt.Errorf("Failed to read CSV")
    }

    return nil
}


func PrintAll() error {

    records, err := ReadCSV("db.csv")
    if err != nil {
       return fmt.Errorf("failed to read CSV: %w", err)
    }

    if len(records) == 1 {
        fmt.Println("No people in db")
        return nil
    }

    for _, record := range records {
        fmt.Println(record)
    }


	return nil
}