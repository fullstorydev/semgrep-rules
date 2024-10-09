package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	queryDb()
	queryDb_fp()
}

func queryDb() {
	db, err := sql.Open("sqlite3", "example.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT field FROM example_table")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		// query fields for each row
		fields, err := db.Query("SELECT ...")
		if err != nil {
			log.Println("Error querying fields:", err)
			continue
		}

		// ruleid: defer-in-loop
		defer fields.Close()

		// Process fields
		var fieldValue string
		if fields.Next() {
			err := fields.Scan(&fieldValue)
			if err != nil {
				log.Println("Error scanning field value:", err)
				continue
			}
			fmt.Println("Field Value:", fieldValue)
		}
	}
}

func queryDb_fp() {
	db, err := sql.Open("sqlite3", "example.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT field FROM example_table")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	// ok: defer-in-loop
	for rows.Next() {
		// query fields for each row
		fields, err := db.Query("SELECT ...")
		if err != nil {
			log.Println("Error querying fields:", err)
			continue
		}

		// Process fields
		var fieldValue string
		if fields.Next() {
			err := fields.Scan(&fieldValue)
			if err != nil {
				log.Println("Error scanning field value:", err)
				continue
			}
			fmt.Println("Field Value:", fieldValue)
		}
		fields.Close()
	}
}
