package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
) //

func ensureDir(path string) error {
	if info, err := os.Stat(path); err == nil {
		if info.IsDir() {
			return nil
		}
		return fmt.Errorf("Path exists but is not a directory: %s", path)
	} else if os.IsNotExist(err) {
		return os.MkdirAll(path, 0755)
	} else {
		return fmt.Errorf("Error checking directory: %w", err)
	}
}
func main() {
	fmt.Println("Checking db directory...")
	err := ensureDir("db")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ensure Dir error: ", err)
		panic(err)
	}

	db, err := sql.Open("sqlite3", "db/main.db")
	if err != nil {
		fmt.Fprintln(os.Stderr, "SQL Open error: ", err)
		panic(err)
	}
	defer db.Close()
	fmt.Println("Opened database")

	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		panic(err)
	}

	var createUsers = `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL UNIQUE,
			hash_password TEXT NOT NULL
		); 
	`

	var createStates = `
		CREATE TABLE IF NOT EXISTS states (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INT NOT NULL,
			name TEXT NOT NULL UNIQUE,
			save TEXT NOT NULL,
    		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

			FOREIGN KEY (user_id) REFERENCES users(id)
		);
	`

	fmt.Println("Executing query...")
	_, err = db.Exec(createUsers + createStates)
	if err != nil {
		fmt.Fprintln(os.Stderr, "SQL Exec error: ", err)
		return
	}

	fmt.Println("Created all needed tables.")
}
