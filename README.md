# `connect` Utility Package

The `connect` package provides a simple way to connect to a MySQL database in Go. It handles database creation (if not already present), establishes a connection, and ensures the database is reachable.

## Features

- Automatically checks if the specified database exists.
- Creates the database if it does not exist.
- Establishes and validates a connection to the MySQL database.

---

## Installation

1. Install the `connect` package by including it in your Go module:

```bash
go get github.com/rudrprasad05/go-sql
```

2. Import the package in your Go project:

```go
import "github.com/rudrprasad05/go-sql/connect"
```

3. Usage

```go
package main

import (
	"log"
	"github.com/rudrprasad05/go-sql/connect"
)

func main() {
	// Database configuration
	config := connect.Config{
		Username: "root",
		Password: "password",
		Host:     "127.0.0.1",
		Port:     3306,
		DbName:   "exampledb",
	}

	// Initialize the database
	db, err := connect.InitDB(&config)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	log.Println("Database connection successful!")
}

```

## Configuration

The Config struct is used to pass database connection parameters:

| Field    | Type   | Description           |
| -------- | ------ | --------------------- |
| Username | string | MySQL server Username |
| Password | string | MySQL server Password |
| Port     | int    | MySQL server port     |
| Host     | string | MySQL server host     |
| DbName   | string | MySQL db name         |
