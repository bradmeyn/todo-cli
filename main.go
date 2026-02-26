package main

import (
	"fmt"
)

func main(){
	TakeInput()
}


func TakeInput() error {
    for {
        fmt.Print("What would you like to do: add, show, delete: ")
        var input string
        _, err := fmt.Scan(&input)
        if err != nil {
            return fmt.Errorf("input error: %w", err)
        }

        switch input {
        case "add":
            
            fmt.Print("Enter the persons details first_name, last_name, email, age: ")

            first, last, email, age := "", "","", 0

            _, err := fmt.Scan(&first, &last, &email, &age)
            if err != nil {
                return fmt.Errorf("input error: %w", err)
            }

            AddPerson(first, last, email, age)
        case "show":
		PrintAll()
        case "delete":
            fmt.Print("Enter the ID of the person to delete: ")

            var input string
            _, err := fmt.Scan(&input)
            if err != nil {
                return fmt.Errorf("input error: %w", err)
            }

            DeletePerson(input)
        default:
            fmt.Println("Invalid command, try again")
        }
    }
}