package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func SaveToCsv(person Person) error {

	// Open file
	file, err := os.OpenFile("db.csv", os.O_APPEND,  0644)

	if err != nil {
		return fmt.Errorf("failed to process data: %w", err)
	}
	defer file.Close()

	// Create a new CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the person data as a new record (slice of strings)
	record := []string {
		person.ID,
		person.FirstName,
		person.LastName,
		strconv.Itoa(person.Age),
		person.Email,
	}

	writer.Write(record)

	return nil

}

func PrintAll() error {

// Open file
	file, err := os.OpenFile("db.csv", os.O_RDONLY,  0644)
	if err != nil {
		return fmt.Errorf("failed to open csv: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("Failed to read file: %w", err)
	}

	for i, record := range records {
		  fmt.Printf("Row %d: %v\n", i, record)
	}
	return nil

}