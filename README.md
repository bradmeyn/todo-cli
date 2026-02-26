# todo-cli

A very simple command-line app in Go for storing people in a CSV file.

## What it does

- Adds a person record
- Shows all saved records
- Deletes a person by ID

Data is stored in `db.csv`.

## Requirements

- Go 1.24+

## Run

From the project root:

```bash
go run .
```

You can also build and run the binary:

```bash
go build -o todo-cli
./todo-cli
```

## Usage

When running, the app prompts:

```text
What would you like to do: add, show, delete:
```

### `add`

Prompts for:

```text
Enter the persons details first_name, last_name, email, age:
```

Enter values separated by spaces, for example:

```text
Jane Doe jane@example.com 31
```

### `show`

Prints all rows currently in `db.csv`.

### `delete`

Prompts for a person ID and removes that record from `db.csv`.

## Project files

- `main.go` - interactive CLI loop
- `people.go` - person operations (`add`, `show`, `delete`)
- `storage.go` - CSV read/write helpers
- `db.csv` - data file
