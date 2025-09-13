package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("MCP Server starts working...")

	connStr := "host=localhost port=5432 user=postgres password=12345 dbname=mydb sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Incorrect connection:", err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Ping error:", err)
		return
	}

	_,err = db.Exec("INSERT INTO users(name,email) VALUES($1,$2)","Ruzanna","ruzs3145@gmail.com")
	if err != nil {
		fmt.Println("Insert error:",err)
	}

	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		fmt.Println("Select error",err)
		return 
	}

	defer rows.Close()

	for rows.Next(){
		var id int
		var name,email string
		err = rows.Scan(&id,&name,&email)
		if err != nil { 
			fmt.Println("Scan error",err)
		}
		fmt.Println(id,name,email)
	}

	fmt.Println("Connected to DB successfully.")
}
