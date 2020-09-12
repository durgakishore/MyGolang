package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
)

var (
	server   = "localhost"
	port     = "1433"
	user     = "rapiscan"
	password = "Anubis2830"
	db       *sql.DB
)

func openDBConnection() {
	//db, err := sql.Open("mysql", "rapiscan:Anubis2830@(localhost:1433)/SysDb?charset=utf8&parseTime=True&loc=Local")

	//connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s", server, user, password, port)
	connString := fmt.Sprintf("rapiscan:Anubis2830@tcp(127.0.0.1:1433)/SysDb")

	// Create connection pool
	db, err := sql.Open("sqlserver", connString)

	if err != nil {
		log.Fatal("Error creating connection pool: " + err.Error())
	}
	log.Printf("Connected!\n")

	defer db.Close()

	// Execute the query
	rows, err := db.Query("SELECT * FROM [SysDb].[dbo].[T_Bag_Info]")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	defer rows.Close()
	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	var rowCount int
	for rows.Next() {
		rowCount++
		fmt.Println("Row number ", rowCount)
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)
		}
		fmt.Println("-----------------------------------")
	}
	rowCount = 0
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}
